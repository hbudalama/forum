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

function prepareFrame(e) {
    const ifrm = document.createElement("iframe");
    ifrm.setAttribute("src", "/posts/1");
    ifrm.style.width = "640px";
    ifrm.style.height = "480px";
    ifrm.style.position = "fixed";
    ifrm.style.top = "50%";
    ifrm.style.left = "50%";
    ifrm.style.transform = "translate(-50%, -50%)";
    ifrm.style.backgroundColor = "white";
    ifrm.style.zIndex = 101;
    ifrm.style.border = "1px solid #ccc"; 
    ifrm.style.boxShadow = "0 4px 8px rgba(0,0,0,0.1)"; 
    document.body.appendChild(ifrm);
}

document.addEventListener('DOMContentLoaded', function () {
    var myPostsCheckbox = document.getElementById('myPosts');
    if (myPostsCheckbox) {
        myPostsCheckbox.addEventListener('change', function () {
            if (this.checked) {
                window.location.href = '/myPosts';
            } else {
                window.location.href = '/';
            }
        });
    }
});
