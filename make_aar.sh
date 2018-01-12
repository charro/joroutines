if [ "$#" -ne 2 ]
  then
    echo "You need to provide the relative path of this lib into your GOPATH and the full path where you want the .aar lib to be generated"
    echo
    message="Example:    "
    message+=$0
    message+=" github.com/charro/joroutines  /home/myself/myproject/joroutine.aar"
    echo $message
    echo
else
    gomobile bind -v -target=android -o=$2 $1
fi
