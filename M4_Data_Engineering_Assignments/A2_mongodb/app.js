const MongoClient = require('mongodb').MongoClient;

const uri = "mongodb://127.0.0.1:27017";

const dbName = 'mydb';
const collectionName = 'customers';
const client = new MongoClient(uri);

console.log('Connecting to MongoDB server...');
async function connectToMongoDB() {
    try {
        await client.connect();
        console.log('Connected to MongoDB server');

        const db = client.db(dbName);
        const collection = db.collection(collectionName);

        const documents = [
            { name: 'Jane Doe' },
            { name: 'Alice' },
            { name: 'Bob' }
        ];

        const query = { name: 'John Doe' };
        const result = await collection.findOne(query);
        console.log('Found a document in the customers collection:', result);

    } catch (error) {
        console.error('Error connecting to MongoDB server', error);
    } finally {
        await client.close();
        console.log('Closed connection to MongoDB server');
    }
}

connectToMongoDB();