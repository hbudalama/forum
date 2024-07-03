const logoutBtn = document.getElementById('logoutBtn');
logoutBtn.addEventListener("click", () => {
    fetch('/logout', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Logout failed');
        }
    })
    .then(data => {
        console.log("Logout response:", data);
        if (data.success) {
            window.location.href = '/login'; // Redirect to login page
        } else {
            alert(data.message || "An error occurred during logout.");
        }
    })
    .catch(error => {
        console.log("Logout error:", error);
        alert("An error occurred: " + error.message);
    });
});