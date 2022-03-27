db.createUser({
    user: "ppnext",
    pwd: "ppnext",
    roles: [{
        role: "readWrite",
        db: "ppnext"
        }
    ]
})

db = new Mongo().getDB("ppnext");

db.createCollection('rooms');
db.rooms.createIndex({ name: 1 });