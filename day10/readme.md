#Solution
1) Builds a graph of all the rules, connecting each high and low output to corresponding child  
2) Sorts all the values  
3) Propagates, for each value, the lowest value down the branches until reaching the output (because the values are sorted we know that there cannot be a lower value for each node). If the low value is already set, we know that the current value is the high value for the node.  