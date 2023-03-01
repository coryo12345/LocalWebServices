# Synchronized Polling Queue
## Development Setup
### Running
Make sure to install all module dependencies from go.mod
Run with 
```bash
go run main.go
```

### Testing
This package uses [github.com/golang/mock](https://github.com/golang/mock) to generate mocks for testing. Make sure to see the install instructions to install the `mockgen` tool.

Before running tests - be sure to run the following: (assuming you are in the queue project directory)
```bash
go generate ./...
```

Then test with:
```bash
go test ./...
```

## About
A synchronized polling queue is a distributed messaging system that provides a reliable way for applications to exchange messages with one another. AWS SQS (Simple Queue Service) is one such service offered by Amazon Web Services. Here is a high-level overview of how a synchronized polling queue like AWS SQS works:

    Producer sends message to the queue:
    A producer application sends a message to the synchronized polling queue by specifying the queue name and message content.

    Queue stores message:
    The message is stored in the queue, and is available for any consumer application to retrieve.

    Consumers poll the queue:
    Consumer applications periodically poll the queue to retrieve messages. When a consumer polls the queue, it retrieves messages in batches of up to 10 messages at a time, or fewer if there are fewer messages in the queue.

    Message visibility:
    When a message is retrieved by a consumer, it becomes invisible to other consumers for a configurable amount of time, called the "visibility timeout." This ensures that only one consumer processes a message at a time.

    Processing message:
    The consumer application processes the message, and when it is done, it deletes the message from the queue.

    Retry:
    If a message is not deleted within the visibility timeout, it becomes visible to other consumers again, allowing another consumer to process the message.

By using a synchronized polling queue like AWS SQS, applications can decouple the components of their architecture and achieve greater scalability, reliability, and fault tolerance.

## Things to do
* Add option to disable synchronization? would allow for "global" events