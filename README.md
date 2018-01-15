# joroutines
Get all the power of Goroutines in Android. Android native lib written in Go that lets you run your Android Java code into a goroutine in a very easy way.

## Motivation
Golang relies very much in goroutines to do its background work, and thanks to that, they are really fast and optimized, specially when compared with Threads or even pools of Threads. 

Talking about Joroutines, I made some tests launching 200.000 concurrent AsyncTasks and same amount of Joroutines, and they use around 30% less memory and more important, **they are completed in less than half of the time** in average.  

## How it works

joroutines uses [gomobile](https://github.com/golang/mobile) for the .aar to be compiled and packaged. Gomobile creates all required Java stubs and the native .so lib that will be called from Java code via JNI.

## How to use it

### A) The easy way: Add the dependency to your gradle file

The .aar file is hosted in JCenter default android repository, so you just need to add this line inside 'dependencies' section of your project and it will be added to your project:
```
compile 'com.github.charro:joroutines:0.0.1'
```

### B) The hard way: Build the aar yourself

Checkout this lib into your GOPATH:

go get github.com/charro/joroutines

Get and configure gomobile in case you haven't done that yet, following the instructions from [Gomobile](https://github.com/golang/go/wiki/Mobile)

Build the .aar file using the make_aar.sh command:

make_aar.sh [RELATIVE_PATH_OF_JOROUTINES_INTO_YOUR_GOPATH] [PATH_AND_NAME_OF_AAR_FILE_TO_GENERATE]

example:  make_aar.sh github.com/charro/joroutines /home/myself/joroutines.aar

The script will call gomobile to do its work, so in case you get some error, you should check that you configured Gomobile correctly (Check if you did the init, set the SDK and NDK ... etc)

If everything went right, you'll have now a new .aar file that you can import in your Android Studio project as a new module. You'll get also a .jar file with the sources in case you need to debug it.


## How to use it in your code

Once you've included the lib in your project you can use the Joroutines Java classes to create your own background tasks like this:

```
Joroutines.runIntoGoroutine(new BackgroundTask(){
            @Override public void run(){
                // Code that will run into a goroutine in background
            }

            @Override public void logDebug(String message) {
                // Message coming from go
                Log.d("GODEBUG", message);
            }
        });
```
