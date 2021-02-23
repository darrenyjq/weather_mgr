cur_path=`pwd`
echo $cur_path
res=`docker ps | grep 'go_dev_env' | awk '{print $(NF-0)}'`
echo $res
name="go_dev_env"
echo $name
if [[ $res == $name ]]
then
    echo 'docker is running, go into>>>>>'
    docker exec -it go_dev_env /bin/sh
else

    docker run --rm=true -dit  -v/Users/Fun/Documents/code/go:/go   -p 9529:80 -p 6379:6379 --name go_dev_env  harbor.cootekservice.com/ime_us/go_dev_env:1.0.1
    docker exec -it go_dev_env /bin/sh

fi