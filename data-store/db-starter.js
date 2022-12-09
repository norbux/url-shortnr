db.seq.insertOne({_id: 'counter', seq: 1});
db.seq.findAndModify(
    {
        query: {_id: 'counter'},
        update: {
            $inc: {seq: 1}
        },
        new: true
    }
);
