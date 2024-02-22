# YetiDB

YetiDB is a Golang project that provides an interface for interacting with a database using the Protocol Buffers data format.

## Usage

To use YetiDB, you need to implement the `Driver` interface defined in `driver.go` for your specific database. Then, you can create a new session, execute transactions, and handle results as per your requirements.

## How to play with Yeti

1. Clone the project

    `git clone git@github.com:sauvikbiswas/yeti.git`

2. Build the protoc plugin

    `make install-yeti-plugin`

3. Compile the test proto files

    `make yeti-proto`

4. Run some tests

    `make yeti-tests`

4. Profit!