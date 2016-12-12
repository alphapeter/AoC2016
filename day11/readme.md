# Algorithm

The state is represented as a single vector, representing the number of generators and microchips on each floor.
<GeneratorsFloor1, MicrochipsFloor1, ..., GeneratorsFloor4, MicrochipsFloor4, ElevatorPosition>

The problem is generalized into simple rules: 
* There cannot be more microchips than generators on a single floor
* If there are no generators, the number of microchip can be larger than the number of generators
* The position of the elevator has to be between 0 and 3
* There cannot be a negative number of equipment on any floor

The moves are, for each adjacent floors 
Represented as vectors.   
<0,-1,0,1,1>  
<-1,0,1,0,1>  
<-1,-1,1,1,1>  
<-2,0,2,0,1>  
<0,-2,2,0,1>  

and 

<0,1,0,-1,-1>    
...  
<0,2,0,-2,-1>  

where <0,-1,1,1> moves one microchip up one level

All possible moves from the current state, in respect to the constraints, are added as edges. This creates a graph of all possible states and the path between them.    
Dijkstras algorithm is used to find the shortest path between the start state and the end state.