if [ $# -eq 0 ]
then
    echo "usage:"
    echo "holistic.sh install - installs npm packages"
    echo "holistic.sh build - builds the client and server"
    echo "holistic.sh start - starts the server"
    echo "holistic.sh stop - stops the server"
    echo "holistic.sh update - updates npm packages"
    echo "holistic.sh clean - removes all build artifacts"
    exit 0
fi


if [ $1 == 'clean' ]
then
cd src/holistic-client
rm -fr build
cd ../..

cd src/holistic-server
rm -f holistic
cd ../..

rm -fr dist/static 
rm -fr dist/holistic
rm -fr dist/server.crt
rm -fr dist/server.key
echo clean successfull
fi

if [ $1 == 'install' ]
then
cd src/holistic-client
npm install
cd ../..
echo install successfull
fi

if [ $1 == 'update' ]
then
cd src/holistic-client
npm update
cd ../..
echo update successfull
fi

if [ $1 == 'build' ]
then
cd src/holistic-client
npm run build
cd ../..
echo build client successfull

rm -fr dist/static
cp -fr src/holistic-client/build dist/static 

cd src/holistic-server
go build
cd ../..

cp -f src/holistic-server/holistic dist/holistic
cp -f src/holistic-server/server.crt dist/server.crt
cp -f src/holistic-server/server.key dist/server.key
echo build server successfull
fi

if [ $1 == 'start' ]
then
cd dist
./holistic &
echo $! > .pid
cd ../
fi

if [ $1 == 'stop' ]
then
cd dist
cat .pid | xargs kill
rm .pid
cd ../
fi