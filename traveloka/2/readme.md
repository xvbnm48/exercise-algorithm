soal 2
the jungle book
There are a number of animal species in the jungle.  Each species has one or more predators that may be direct or indirect.   Species X is said to be a predator of species Y if at least one of the following is true:

species X is a direct predator of species Y.
if species X is a direct predator of species Z, and Z is a direct predator of Y, then species X is an indirect predator of species Y. Indirect predation is transitive through any number of levels.

each species has a maximum of 1 direct predator. No two species will ever be mutual predators, and no species is a predator of itself. Determine the minimum number of gorupus that must be formed to so that no species is grouped with its predatoirs, direct or indirect.

exmaple

predators = [-1, 8, 6, 0, 7, 3, 8, 9, -1, 6]

Each position in predators represents a species and each element represents a predator of that species, or -1 if there are none. The graph of predation is below using zero based indexing. All labels are the indices within predators

From the graph, a possible grouping is:
[0,8]
[3,1,6]
[5,2,9]
[7]
[4]

A minimum of  5  groups are needed to satisfy all conditions.

minimumGroups has the following parameter(s):
int predators[n]:  an array of integers where predator[i]  represents the direct predator of the ith
species or -1 if there is none.

returns:
int the minimum number of groups formed given the rule that none of the species will associate with its predators