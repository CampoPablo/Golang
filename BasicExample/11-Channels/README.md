# Channels

In Golang, or Go, channels are a means through which different goroutines communicate.

Think of them as pipes through which you can connect with different concurrent goroutines. The communication is bidirectional by default, meaning that you can send and receive values from the same channel.

Moreover, by default, channels send and receive until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

