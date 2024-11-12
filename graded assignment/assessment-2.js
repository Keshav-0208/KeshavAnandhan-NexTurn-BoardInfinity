// --------------------------------------- part 1 ----------------------------------------------------- //

//1.Create the Collections and Insert Data:

db.customers.insertMany([
    { 
        _id: ObjectId("64a7f9b6a4b9f07a6b010101"), 
        name: "Michael De Santa", 
        email: "michael.desanta@gta.com", 
        address: { street: "3671 Whispymound Dr", city: "Los Santos", zipcode: "10001" }, 
        phone: "555-0101", 
        registration_date: new Date("2023-01-01T12:00:00Z") 
    },
    { 
        _id: ObjectId("64a7f9b6a4b9f07a6b010102"), 
        name: "Franklin Clinton", 
        email: "franklin.clinton@gta.com", 
        address: { street: "3671 Forum Dr", city: "Los Santos", zipcode: "10002" }, 
        phone: "555-0202", 
        registration_date: new Date("2023-02-15T12:00:00Z") 
    },
    { 
        _id: ObjectId("64a7f9b6a4b9f07a6b010103"), 
        name: "Trevor Philips", 
        email: "trevor.philips@gta.com", 
        address: { street: "Ron Alternates Wind Farm", city: "Blaine County", zipcode: "10003" }, 
        phone: "555-0303", 
        registration_date: new Date("2023-03-10T12:00:00Z") 
    },
    { 
        _id: ObjectId("64a7f9b6a4b9f07a6b010104"), 
        name: "Carl Johnson", 
        email: "cj.johnson@gta.com", 
        address: { street: "Ganton St", city: "Los Santos", zipcode: "10006" }, 
        phone: "555-0606", 
        registration_date: new Date("2023-04-20T12:00:00Z") 
    },
    { 
        _id: ObjectId("64a7f9b6a4b9f07a6b010105"), 
        name: "Tommy Vercetti", 
        email: "tommy.vercetti@gta.com", 
        address: { street: "Ocean View", city: "Vice City", zipcode: "10007" }, 
        phone: "555-0707", 
        registration_date: new Date("2023-05-05T12:00:00Z") }]);


db.orders.insertOne({ 
    "_id": ObjectId(), 
    "order_id": "ORD123456", 
    "customer_id": ObjectId("64a7f9b6a4b9f07a6b010101"), 
    "order_date": new Date("2023-05-15T14:00:00Z"), 
    "status": "shipped", 
    "items": [{ "product_name": "Laptop", "quantity": 1, "price": 1500 }, { "product_name": "Mouse", "quantity": 2, "price": 25 }], 
    "total_value": 1550 });

db.orders.insertOne({ 
    "_id": ObjectId(), 
    "order_id": "ORD234567", 
    "customer_id": ObjectId("64a7f9b6a4b9f07a6b010102"), 
    "order_date": new Date("2023-06-10T10:00:00Z"), 
    "status": "delivered", 
    "items": [{ "product_name": "Smartphone", "quantity": 1, "price": 800 }], 
    "total_value": 800 });

db.orders.insertOne({ 
    "_id": ObjectId(), 
    "order_id": "ORD345678", 
    "customer_id": ObjectId("64a7f9b6a4b9f07a6b010103"), 
    "order_date": new Date("2023-07-01T16:30:00Z"), 
    "status": "pending", 
    "items": [{ "product_name": "Tablet", "quantity": 1, "price": 600 }], 
    "total_value": 600 });

db.orders.insertOne({ 
    "_id": ObjectId(), 
    "order_id": "ORD456789", 
    "customer_id": ObjectId("64a7f9b6a4b9f07a6b010104"), 
    "order_date": new Date("2023-08-15T12:45:00Z"), 
    "status": "shipped", 
    "items": [{ "product_name": "Headphones", "quantity": 3, "price": 50 }], 
    "total_value": 150 });

db.orders.insertOne({ 
    "_id": ObjectId(), 
    "order_id": "ORD567890", 
    "customer_id": ObjectId("64a7f9b6a4b9f07a6b010105"), 
    "order_date": new Date("2023-09-20T09:00:00Z"), 
    "status": "delivered", 
    "items": [{ "product_name": "Smartwatch", "quantity": 1, "price": 200 }, { "product_name": "Desk Lamp", "quantity": 1, "price": 35 }], 
    "total_value": 235 });


//2.Find Orders for a Specific Customer:

db.orders.find({ customer_id: ObjectId("64a7f9b6a4b9f07a6b010101") });

//output

/*
[
    {
      _id: ObjectId('6732f50e66b76ace680d8191'),
      order_id: 'ORD123456',
      customer_id: ObjectId('64a7f9b6a4b9f07a6b010101'),
      order_date: ISODate('2023-05-15T14:00:00.000Z'),
      status: 'shipped',
      items: [
        { product_name: 'Laptop', quantity: 1, price: 1500 },
        { product_name: 'Mouse', quantity: 2, price: 25 }
      ],
      total_value: 1550
    }
  ]
*/

//3.Find the Customer for a Specific Order:

db.customers.findOne({ "_id": db.orders.findOne({ order_id: "ORD123456" }).customer_id });

//output

/*
{
  _id: ObjectId('64a7f9b6a4b9f07a6b010101'),
  name: 'Michael De Santa',
  email: 'michael.desanta@gta.com',
  address: {
    street: '3671 Whispymound Dr',
    city: 'Los Santos',
    zipcode: '10001'
  },
  phone: '555-0101',
  registration_date: ISODate('2023-01-01T12:00:00.000Z')
}
*/

//4.Update Order Status:

db.orders.updateOne({ order_id: "ORD123456" }, { $set: { status: "delivered" } });

//output

/*
{
  acknowledged: true,
  insertedId: null,
  matchedCount: 1,
  modifiedCount: 1,
  upsertedCount: 0
}
*/

//5. Delete an Order:

db.orders.deleteOne({ order_id: "ORD123456" });

//Output 

/*
{ acknowledged: true, deletedCount: 1 }
 */

// --------------------------------------- part 2 ----------------------------------------------------- //

//1. Calculate Total Value of All Orders by Customer:

db.orders.aggregate([
    { $group: { _id: "$customer_id", total_value: { $sum: "$total_value" } } },
    { $lookup: { from: "customers", localField: "_id", foreignField: "_id", as: "customer_details" } },
    { $unwind: "$customer_details" },
    { $project: { _id: 0, customer_name: "$customer_details.name", total_order_value: "$total_value" } }]);

//output

/*
[
  { customer_name: 'Franklin Clinton', total_order_value: 800 },
  { customer_name: 'Trevor Philips', total_order_value: 600 },
  { customer_name: 'Tommy Vercetti', total_order_value: 235 },
  { customer_name: 'Carl Johnson', total_order_value: 150 }
]
*/

//2. Group Orders by Status:

db.orders.aggregate([
    { $group: { _id: "$status", order_count: { $sum: 1 } } }]);

//output

/*
[
  { _id: 'shipped', order_count: 1 },
  { _id: 'delivered', order_count: 2 },
  { _id: 'pending', order_count: 1 }
]
 */

//3. List Customers with Their Recent Orders:

db.orders.aggregate([
    { $sort: { order_date: -1 } },
    { $group: { _id: "$customer_id", most_recent_order: { $first: "$$ROOT" } } },
    { $lookup: { from: "customers", localField: "_id", foreignField: "_id", as: "customer_details" } },
    { $unwind: "$customer_details" },
    { $project: { _id: 0, customer_name: "$customer_details.name", email: "$customer_details.email", order_id: "$most_recent_order.order_id", total_value: "$most_recent_order.total_value", order_date: "$most_recent_order.order_date" } }]);

//output

/*
[
  {
    customer_name: 'Carl Johnson',
    email: 'cj.johnson@gta.com',
    order_id: 'ORD456789',
    total_value: 150,
    order_date: ISODate('2023-08-15T12:45:00.000Z')
  },
  {
    customer_name: 'Tommy Vercetti',
    email: 'tommy.vercetti@gta.com',
    order_id: 'ORD567890',
    total_value: 235,
    order_date: ISODate('2023-09-20T09:00:00.000Z')
  },
  {
    customer_name: 'Trevor Philips',
    email: 'trevor.philips@gta.com',
    order_id: 'ORD345678',
    total_value: 600,
    order_date: ISODate('2023-07-01T16:30:00.000Z')
  },
  {
    customer_name: 'Franklin Clinton',
    email: 'franklin.clinton@gta.com',
    order_id: 'ORD234567',
    total_value: 800,
    order_date: ISODate('2023-06-10T10:00:00.000Z')
  }
]
*/

//4. Find the Most Expensive Order by Customer:

db.orders.aggregate([
    { $sort: { total_value: -1 } },
    { $group: { _id: "$customer_id", most_expensive_order: { $first: "$$ROOT" } } },
    { $lookup: { from: "customers", localField: "_id", foreignField: "_id", as: "customer_details" } },
    { $unwind: "$customer_details" },
    { $project: { _id: 0, customer_name: "$customer_details.name", order_id: "$most_expensive_order.order_id", total_value: "$most_expensive_order.total_value" } }]);

//output

/*
[
  {
    customer_name: 'Tommy Vercetti',
    order_id: 'ORD567890',
    total_value: 235
  },
  {
    customer_name: 'Franklin Clinton',
    order_id: 'ORD234567',
    total_value: 800
  },
  {
    customer_name: 'Carl Johnson',
    order_id: 'ORD456789',
    total_value: 150
  },
  {
    customer_name: 'Trevor Philips',
    order_id: 'ORD345678',
    total_value: 600
  }
]
*/

// --------------------------------------- part 3 ----------------------------------------------------- //

//1. Find All Customers Who Placed Orders in the Last Month:

db.orders.aggregate([
    { $match: { order_date: { $gte: new Date(new Date().setMonth(new Date().getMonth() - 1)) } } },
    { $lookup: { from: "customers", localField: "customer_id", foreignField: "_id", as: "customer_details" } },
    { $unwind: "$customer_details" },
    { $project: { _id: 0, customer_name: "$customer_details.name", order_id: "$order_id", total_value: "$total_value", order_date: "$order_date" } }]);

// --------------------------------------- part 4 ----------------------------------------------------- //

//1. Find Customers Who Have Not Placed Orders:

db.customers.aggregate([
    { $lookup: { from: "orders", localField: "_id", foreignField: "customer_id", as: "orders" } },
    { $match: { orders: { $size: 0 } } },
    { $project: { _id: 0, name: 1, email: 1 } }]);

//output

/*
[ { name: 'Michael De Santa', email: 'michael.desanta@gta.com' } ]
*/

//2. Calculate the Average Number of Items Ordered per Order:

db.orders.aggregate([
    { $unwind: "$items" },
    { $group: { _id: "$_id", items_ordered: { $sum: "$items.quantity" } } },
    { $group: { _id: null, average_items: { $avg: "$items_ordered" } } }]);

//output

/*
[ { _id: null, average_items: 1.75 } ]
*/

//3. Join Customer and Order Data Using $lookup:

db.orders.aggregate([
    { $lookup: { from: "customers", localField: "customer_id", foreignField: "_id", as: "customer_details" } },
    { $unwind: "$customer_details" },
    { $project: { _id: 0, customer_name: "$customer_details.name", email: "$customer_details.email", order_id: 1, total_value: 1, order_date: 1 } }]);

//output

/*
[
  {
    order_id: 'ORD234567',
    order_date: ISODate('2023-06-10T10:00:00.000Z'),
    total_value: 800,
    customer_name: 'Franklin Clinton',
    email: 'franklin.clinton@gta.com'
  },
  {
    order_id: 'ORD345678',
    order_date: ISODate('2023-07-01T16:30:00.000Z'),
    total_value: 600,
    customer_name: 'Trevor Philips',
    email: 'trevor.philips@gta.com'
  },
  {
    order_id: 'ORD456789',
    order_date: ISODate('2023-08-15T12:45:00.000Z'),
    total_value: 150,
    customer_name: 'Carl Johnson',
    email: 'cj.johnson@gta.com'
  },
  {
    order_id: 'ORD567890',
    order_date: ISODate('2023-09-20T09:00:00.000Z'),
    total_value: 235,
    customer_name: 'Tommy Vercetti',
    email: 'tommy.vercetti@gta.com'
  }
]
*/



