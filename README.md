# badmintonClub

my ThoughtWork homework(repository [https://github.com/shamaoxiaogui/badmintonClub](https://github.com/shamaoxiaogui/badmintonClub))

## Install

### from github

1. Install Golang environment, **be sure to set $GOPATH correctly**, read the instructions [here](https://golang.org/doc/install)
1. clone the project:

    ```
    go get github.com/shamaoxiaogui/badmintonClub
    ```
1. run the unit test:

    ```
    cd $GOPATH/src/github.com/shamaoxiaogui/badmintonClub
    # this step might show you some runtime error, ignore those 
    # because there are illegal input test in the unit test, focus on the last line in stdout,
    # and check if there is a "PASS"
    go test
    # test the sub lib
    cd activity
    go test
    cd ../strategy
    go test
    cd ..
    ```
1. run the example:

    ```
    cd $GOPATH/src/github.com/shamaoxiaogui/badmintonClub/example
    go build
    # the demoDatas file have all example input mentioned in the pdf
    ./example < demoDatas
    ```

### from local file

move the whole project to your ```$GOPATH/src/``` ,so the workspace looks like:

```shell
$GOPATH/src/github.com/shamaoxiaogui/badmintonClub/example
```

then you can run test or example just like section **from github** said

## Usage

Please check the example/example.go

## Summary

The project divides the homework into two parts:

1. struct activity, to hold one record. It can parse the legal input string, and format the output
2. struct strategy, to calculate the income and payment in each record. It can be modified to expand the function

Firstly, user creats a strategy struct by calling ```NewStrategy()``` factory function, so when we modify or add a new strategy, users don't have to modify their code.
Secondly, user creats a activity struct by calling ```NewActivity()``` construct function. It can force the user injects a strategj struct into activity.
And then, in the for loop, use the same activity parse input record, calculate and save the result into an output string. In this way, we can slightly reduce the struct creat cost

## Confusion(ignore these blabla)

I'm not good at high-level programming, so there is my confusion:

1. I'm not sure whether the strategy patten should be used, and also I don't know if I did it correctly .At firt I put everything in activity struct, and for this project it seems not so ugly... Anyway, I move the calculate function out and leave a callback interface in the activity struct then, in other words, change to use functional programming. And finally I decide to use strategy pattern to handle the extendibility
2. I'm not sure if I did the right thing about illegal inputs . The sub libary will ```panic()``` when they meet an illegal input, and the ```GenerateSummary()``` handle the panic and make the output string empty, But it will put error message in stdout
3. As the coding style from google, the struct method's reciver should be a struct instance instead of a struct pointer when the method don't modify the struct. That means every call on these method, the struct be copied, even the method is a getter method. Isn't that too expensive? I'm not good at golang, so I'm really confused...
4. As I said, I'm not good at golang. I'm more familiar with C++ and C(God bless me). 
