const logoutBtn = document.getElementById('logoutBtn');
const postBtn = document.getElementById('post-button-container');

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

postBtn.addEventListener()


