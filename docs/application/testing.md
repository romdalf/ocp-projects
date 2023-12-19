

The project/namespace construct can be leverage in a shared platform to address different needs from dev to production, including test activities. For the latter, we might want to differentiate functional and nonfunctional testing:  
* the functional testing will focus on the application itself, this will consume a certain amount of compute resources and can be forecasted.
* the nonfunctional testing will focus on the resilience in regards of fault, injection, day 2 operations, and more.

The nonfunctional testing might required to put in place a set of dedicated worker nodes to which only dev and test applications can be deployed and on which these scenarios can be ran. No production application should be allowed to be scheduled on this node to avoid unexpected faults. 
