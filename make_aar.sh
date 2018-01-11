if [ -z "$1" ]
  then
    printf "\nYou need to provide the full path where you want the .aar file to be generated. A jar file with sources will be generated there too\n\n"
    message="Example:    "
    message+=$0
    message+=" /home/myself/myproject/joroutine.aar"
    echo $message
    echo
else
  export GOPATH=/home/maik/development/go
  gomobile bind -v -target=android -o=$1 joroutines
fi
