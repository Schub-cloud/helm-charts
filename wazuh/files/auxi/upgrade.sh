#!/bin/bash

# Run the initial command and store the output in a variable
output=$(/var/ossec/bin/agent_upgrade -l)

# Use a while loop to read each line of the output
echo "$output" | while read -r line; do
  # Check if the line contains an agent ID (we assume IDs are three digits here)
  if [[ $line =~ ^[0-9]{3} ]]; then
    # Extract the ID from the line
    id=$(echo "$line" | awk '{print $1}')
    
    # Run the upgrade command for the extracted ID
    /var/ossec/bin/agent_upgrade -a "$id"
  fi
done
