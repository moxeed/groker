# groker

## broker

---

broker **(broker.go-broker/core.go)** service **(broker/service.go)**, is simply a fifo queue. each message regardless of its sync or aysnc nature are added to queue and will be published when became available at the end of queue.

broker suports two types of messages aysnc and sync. is sync scenario client request will be blocked until message reaches end of queue, then the request will be realesed. this mechanism is implemented using mutex lock service will wait until pubish confirmation is triggred (mutex is unlocked).

is aync version once broker sends acceptance response as soon as message message is on the queue. when message reaches end of queue broker sends a publish notification to server. to enrich broker with notification capability, server has to give a webhook to broker for further notification (is also could be done with frequent pulling from broker but push notifications provide low latency)

becaues this methods only defer in communication with server and use the same infrastructure, broker uses sender abstraction AsyncSender **(groker/async_sender.go)** and SyncSender **(groker/sync_sender.go)** are implementations for sender interface.

if no client is available broker will hold messages until at least one client is available. then starts to send messages. **(core.go/publish)**

## server

---

server **(server/server.go)** is simply a load generator. an infinite loop sends message each one second to broker. some messages are async so server should provide webhook for broker to send notifcation of ackknowledgements.

## client

---

client **(client/client.go)** subscribes to broker and provides a webhook for broker to send new message notifcations and prints body of recived messagae

any number of client could be added to broker and it will notify them all about new messages.   
1 - messages before join will not be sent   
2 - broker do not ackknowledge delivery on client side so there is no gurantiy for client to recive every message

each time a client is not accessible by broker it will be removed

## Question Queue Advantages

---

1 - loosly coupling request service and process service thus decreasing complexity (if pointers are used in shared memory apllications need to know about object structers in other application)    
2 - easy scale up : at any time more processors could be atached to queue and increase through put of service (shared memory has concurency problems and need more aware design)   
3 - no need to keep all services at same server so it makes creating distributed systems easier (which is a problem in shared memory because we have to spread memory in multi servers)

