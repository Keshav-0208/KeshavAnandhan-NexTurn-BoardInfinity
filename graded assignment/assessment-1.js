let productsJSON = `[
    {"id": 1, "name": "PlayStation 5", "category": "Electronics", "price": 499, "available": true},
    {"id": 2, "name": "Tea Table", "category": "Furniture", "price": 75, "available": true},
    {"id": 3, "name": "Xbox", "category": "Electronics", "price": 450, "available": false},
    {"id": 4, "name": "Blender", "category": "Kitchen", "price": 60, "available": true}
]`;

function parseproducts() {
    return JSON.parse(productsJSON);
}

function addproduct(newProduct) {
    let products = parseproducts();
    products.push(newProduct);
    console.log("Product added:", newProduct);
    return products;
}

function updateproductprice(productId, newPrice) {
    let products = parseproducts();
    let product = products.find(p => p.id === productId);
    if (product) {
        product.price = newPrice;
        console.log(`Price updated for product ID ${productId}: $${newPrice}`);
    } else {
        console.log(`Product with ID ${productId} is not found.`);
    }
    return products;
}

function getavailableproducts() {
    let products = parseproducts();
    return products.filter(p => p.available);
}

function getproductsbycategory(category) {
    let products = parseproducts();
    return products.filter(p => p.category === category);
}

//1. Parse the JSON data:
console.log("Parsed Products:", parseproducts());

//2. Add a new product:
console.log("Adding a new product...");
addproduct({"id": 5, "name": "Wireless Earbuds", "category": "Electronics", "price": 49.99, "available": true});

//3. Update the price of a product:
console.log("Updating product price...");
updateproductprice(1, 520);

//4. Filter products based on availability:
console.log("Available Products:", getavailableproducts());

//5. Filter products by category:
console.log("Products in Electronics Category:", getproductsbycategory("Electronics"));
