We have some clickstream data that we gathered on our client's website. 
Using cookies, we collected snippets of users' anonymized URL histories 
while they browsed the site. The histories are in chronological order, and 
no URL was visited more than once per person.

Write a function that takes two users' browsing histories as input and 
returns the longest contiguous sequence of URLs that appears in both.

Sample input:
user0 = ["/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"]
user1 = ["/start", "/pink", "/register", "/orange", "/red", "a"]
user2 = ["a", "/one", "/two"]
user3 = ["/pink", "/orange", "/yellow", "/plum", "/blue", "/tan", "/red", "/amber", "/CornflowerBlue", "/LightGoldenRodYellow", ]
user4 = ["/pink", "/orange", "/amber", "/plum", "/blue", "/tan", "/red", "/lavender", "/CornflowerBlue", "/LightGoldenRodYellow"]
user5 = ["a"]

Sample output:
findLongestHistory(user0, user1) => ["/pink", "/register", "/orange"]
findLongestHistory(user0, user2) => [] (empty)
findLongestHistory(user0, user0) => ["/start", "/green", "/blue", "/pink", "/register", "/orange", "/one/two"]
findLongestHistory(user5, user2) => ["a"]
findLongestHistory(user3, user4) => ["/plum", "/blue", "/tan", "/red"]
findLongestHistory(user4, user3) => ["/plum", "/blue", "/tan", "/red"]