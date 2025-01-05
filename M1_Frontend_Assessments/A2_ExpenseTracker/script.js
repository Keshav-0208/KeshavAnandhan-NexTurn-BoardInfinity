document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("expense-form");
    const tableBody = document.getElementById("expense-table-body");
    const summary = document.getElementById("category-summary");
    const chartCanvas = document.getElementById("expense-chart");
    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
  
    const saveExpenses = () => {
      localStorage.setItem("expenses", JSON.stringify(expenses));
    };
  
    const updateUI = () => {
      tableBody.innerHTML = "";
      summary.innerHTML = "";
      const totals = {};
  
      expenses.forEach((e, i) => {
        const row = document.createElement("tr");
        row.innerHTML = `
          <td>${e.amount}</td>
          <td>${e.description}</td>
          <td>${e.category}</td>
          <td><button data-index="${i}">Delete</button></td>
        `;
        tableBody.appendChild(row);
  
        if (!totals[e.category]) {
          totals[e.category] = 0;
        }
        totals[e.category] += e.amount;
      });
  
      for (let category in totals) {
        const div = document.createElement("div");
        div.textContent = `${category}: $${totals[category]}`;
        summary.appendChild(div);
      }
  
      if (Object.keys(totals).length > 0) {
        const ctx = chartCanvas.getContext("2d");
        new Chart(ctx, {
          type: "pie",
          data: {
            labels: Object.keys(totals),
            datasets: [{
              data: Object.values(totals),
              backgroundColor: ["#ff6384", "#36a2eb", "#cc65fe", "#ffce56"]
            }]
          }
        });
      }
    };
  
    form.addEventListener("submit", (e) => {
      e.preventDefault();
      const amount = parseFloat(document.getElementById("amount").value);
      const description = document.getElementById("description").value;
      const category = document.getElementById("category").value;
  
      expenses.push({ amount, description, category });
      saveExpenses();
      updateUI();
      form.reset();
    });
  
    tableBody.addEventListener("click", (e) => {
      if (e.target.tagName === "BUTTON") {
        const index = e.target.getAttribute("data-index");
        expenses.splice(index, 1);
        saveExpenses();
        updateUI();
      }
    });
  
    updateUI();
  });
  