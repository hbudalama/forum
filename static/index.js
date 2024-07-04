function logoutHandler(e){
    fetch('/logout', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then(response => {
            window.location.href = ('/login')
        })

}