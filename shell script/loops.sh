#!/bin/sh
INPUT_STRING=hello
while [ "$INPUT_STRING" != "bye" ]
do
  echo "Please type something in (bye to quit)"
  read INPUT_STRING
  echo "You typed: $INPUT_STRING"
done

# for i in hello 1 * 2 goodbye 
# do
#   echo "Looping ... i is set to $i"
# done


# for i in 1 2 3 4 5
# do
#   echo "Looping ... number $i"
# done