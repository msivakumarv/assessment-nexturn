document.addEventListener('DOMContentLoaded', () => {
  const expenseForm = document.getElementById('expense-form');
  const amountInput = document.getElementById('amount');
  const descriptionInput = document.getElementById('description');
  const categoryInput = document.getElementById('category');
  const expenseTableBody = document.getElementById('expense-table-body');
  const chartCanvas = document.getElementById('expenseChart');

  let expenses = JSON.parse(localStorage.getItem('expenses')) || [];
  let expenseChart = null;

  function updateLocalStorage() {
      localStorage.setItem('expenses', JSON.stringify(expenses));
  }

  function renderExpenses() {
      expenseTableBody.innerHTML = '';

      expenses.forEach((expense, index) => {
          const row = document.createElement('tr');
          row.innerHTML = `
              <td>$${expense.amount.toFixed(2)}</td>
              <td>${expense.description}</td>
              <td>${expense.category}</td>
              <td><button class="delete" data-index="${index}">Delete</button></td>
          `;
          expenseTableBody.appendChild(row);
      });

      updateChart();
  }

  function updateChart() {
    const categoryTotals = expenses.reduce((totals, expense) => {
        if (!totals[expense.category]) {
            totals[expense.category] = 0;
        }
        totals[expense.category] += expense.amount;
        return totals;
    }, {});

    const categories = Object.keys(categoryTotals);
    const amounts = Object.values(categoryTotals);

    
    const predefinedColors = [
        '#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF', '#FF9F40'
    ];

    const dynamicColors = [...predefinedColors];
    while (dynamicColors.length < categories.length) {
        dynamicColors.push(generateRandomColor());
    }

    
    if (expenseChart) {
        expenseChart.destroy();
    }

    
    expenseChart = new Chart(chartCanvas, {
        type: 'pie', 
        data: {
            labels: categories,
            datasets: [{
                label: 'Spending by Category',
                data: amounts,
                backgroundColor: dynamicColors,
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {
                    position: 'top',
                },
                tooltip: {
                    callbacks: {
                        label: function(tooltipItem) {
                            return `$${tooltipItem.raw.toFixed(2)}`;
                        }
                    }
                }
            }
        }
    });
}


function generateRandomColor() {
    const letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}


  expenseForm.addEventListener('submit', (event) => {
      event.preventDefault();

      const amount = parseFloat(amountInput.value);
      const description = descriptionInput.value.trim();
      const category = categoryInput.value;

      if (isNaN(amount) || description === '' || category === '') {
          alert('Please fill in all fields correctly!');
          return;
      }

      expenses.push({ amount, description, category });
      updateLocalStorage();
      renderExpenses();

      amountInput.value = '';
      descriptionInput.value = '';
      categoryInput.value = '';
  });

  expenseTableBody.addEventListener('click', (event) => {
      if (event.target.classList.contains('delete')) {
          const index = event.target.dataset.index;
          expenses.splice(index, 1);
          updateLocalStorage();
          renderExpenses();
      }
  });

  
  renderExpenses();
});
