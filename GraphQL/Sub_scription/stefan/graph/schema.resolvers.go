package graph

import (
	"context"
	"log"
	"stefan/graph/model"
	"strconv"
)

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, user string, content string) (string, error) {
	// Construct the newly sent message and append it to the existing messages
	x := len(r.ChatMessages)
	msg := model.Message{
		ID:      strconv.Itoa(x),
		User:    user,
		Content: content,
	}
	r.ChatMessages = append(r.ChatMessages, &msg)

	r.mu.Lock()
	// Notify all active subscriptions that a new message has been posted by posted. In this case we push the now
	// updated ChatMessages array to all clients that care about it.
	for _, observer := range r.ChatObservers {
		observer <- r.ChatMessages
	}
	r.mu.Unlock()

	log.Println("you have succesfully sent a message")
	return msg.ID, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return r.ChatMessages, nil
}

func (r *subscriptionResolver) Messages(ctx context.Context) (<-chan []*model.Message, error) {
	// Create an ID and channel for each active subscription. We will push changes into this channel.
	// When a new subscription is created by the client, this resolver will fire first.
	id := randString(8)
	msgs := make(chan []*model.Message, 1)

	// Start a goroutine to allow for cleaning up subscriptions that are disconnected.
	// This go routine will only get past Done() when a client terminates the subscription. This allows us
	// to only then remove the reference from the list of ChatObservers since it is no longer needed.
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.ChatObservers, id)
		r.mu.Unlock()
	}()
	r.mu.Lock()
	// Keep a reference of the channel so that we can push changes into it when new messages are posted.
	r.ChatObservers[id] = msgs
	r.mu.Unlock()
	// This is optional, and this allows newly subscribed clients to get a list of all the messages that have been
	// posted so far. Upon subscribing the client will be pushed the messages once, further changes are handled
	// in the PostMessage mutation.
	r.ChatObservers[id] <- r.ChatMessages
	log.Println("Successfully called subscription")
	return msgs, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
