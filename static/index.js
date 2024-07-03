const logoutBtn = document.getElementById('logoutBtn');
logoutBtn.addEventListener("click", () => {
    fetch('/logout', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => {
    window.location.href = ('/login')
    })
 
});