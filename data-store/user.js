db.createUser(
    {
        user:"mongoUser",
        pwd:"1234",
        roles: [ 
            { role: "readWrite", db: "seq" }
        ]
    }
);