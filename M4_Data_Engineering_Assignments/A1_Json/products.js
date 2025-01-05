const productData = `[
    { "itemID": 1, "itemName": "PS5", "itemType": "Gaming Console", "cost": 500, "inStock": true },
    { "itemID": 2, "itemName": "Xbox Series S", "itemType": "Gaming Console", "cost": 300, "inStock": false },
    { "itemID": 3, "itemName": "Cyberpunk 2077", "itemType": "Video Game", "cost": 60, "inStock": true },
    { "itemID": 4, "itemName": "Gaming Mouse", "itemType": "Accessories", "cost": 50, "inStock": true },
    { "itemID": 5, "itemName": "Gaming Chair", "itemType": "Furniture", "cost": 250, "inStock": false }
  ]`;
  
  let inventory = JSON.parse(productData);
  
  function addNewItem(itemID, itemName, itemType, cost, inStock) {
    const newItem = { itemID, itemName, itemType, cost, inStock };
    inventory.push(newItem);
    console.log('New item added to inventory:', JSON.stringify(newItem, null, 2));
  }
  
  function modifyItemCost(itemID, updatedCost) {
    const item = inventory.find(i => i.itemID === itemID);
    if (item) {
      item.cost = updatedCost;
      console.log(`Cost for item ID ${itemID} updated to $${updatedCost}`);
    } 
    else
     {
      console.log('Error: Item not found!');
    }
  }
  
  function showStockItems() {
    const stockAvailable = inventory.filter(i => i.inStock);
    console.log('Available items in stock:');
    stockAvailable.forEach(item => {  console.log(JSON.stringify(item, null, 2));});
  }
  
  function listItemsByCategory(itemType) {
    const itemsByCategory = inventory.filter(i => i.itemType === itemType);
    console.log(`Items in category "${itemType}":`);
    itemsByCategory.forEach(item => {
      console.log(JSON.stringify(item, null, 2));
    });
  }
  
  addNewItem(6, 'Nintendo Switch', 'Gaming Console', 350, true);
  modifyItemCost(3, 40);
  showStockItems();
  listItemsByCategory('Gaming Console');
  