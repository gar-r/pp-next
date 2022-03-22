db.createUser({
    user: "ppnext",
    pwd: "ppnext",
    roles: [{
        role: "readWrite",
        db: "ppnextdb"
        }
    ]
})

db = new Mongo().getDB("ppnexdb");

db.createCollection('rooms');
