# Project Title

## Table of Contents

- [About](#about)
- [How to build](#build)
- [Usage](#usage)

## About <a name = "about"></a>

The repository contains solution for candystore assignment. It fetches input data from URL provided in env file returns to customers. 


### Dependencies

The program is built on go 1.8.3 version. It uses a few third party go libraries which can be installed as below

```
go get github.com/sirupsen/logrus
go get github.com/spf13/viper
go get github.com/PuerkitoBio/goquery
```

### How to build <a name = "build"></a>

Take clone of the repository.

```
git clone https://github.com/rhishimagdum/candystore.git
```

Build

```
cd candystore
go build
```

Run unit test

```
go test entity/customer.go entity/customer_test.go -v
go test config/config.go config/config_test.go -v
```

## Usage <a name = "usage"></a>

Run the binary 

```
./candystore
```
