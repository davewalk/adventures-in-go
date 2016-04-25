# graceful-shutdown

This is an experiment in gracefully shutting down a process that reads from a 
Kafka consumer when an `SIGINT`, `SIGQUIT` or `SIGTERM` is sent to the process.  

When the signal is received, the process will close the Kafka consumer and wait 
for the last message to be processed (just a five second sleep to simluate) before
 exiting.
