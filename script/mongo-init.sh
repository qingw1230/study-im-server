set -e

mongosh <<EOF
use admin
db.auth('root', 'qin1002')
db = db.getSiblingDB('studyIM')
EOF