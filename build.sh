rm -f c2s debug

if [ -z ${ORACLE_SID+x} ]; then
   echo "Sourcing db.env"
   . ~/db.env
else 
   echo "Oracle is setup"
fi

go build .

