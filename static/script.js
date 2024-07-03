document.addEventListener('DOMContentLoaded', () => {
    const loginBtn = document.getElementById('loginBtn');
    const registerBtn = document.getElementById('registerBtn');
    const loginField = document.getElementById('loginField');
    const registerField = document.getElementById('registerField');
    const loginError = document.getElementById('loginError');
    const registerError = document.getElementById('registerError');
    loginBtn.addEventListener('change', () => {
        loginField.style.display = 'flex';
        registerField.style.display = 'none';
        loginError.style.display = 'none';
        registerError.style.display = 'none';
    });
    registerBtn.addEventListener('change', () => {
        loginField.style.display = 'none';
        registerField.style.display = 'flex';
        loginError.style.display = 'none';
        registerError.style.display = 'none';
    });
    loginField.addEventListener('submit', (e) => {
        e.preventDefault();
        const username = loginField.querySelector('input[name="username"]').value;
        const password = loginField.querySelector('input[name="password"]').value;
        loginError.style.display = 'none';
        if (username.includes(' ')) {
            loginError.textContent = "Username cannot contain spaces.";
            loginError.style.display = 'block';
            return;
        }
        if (password.length < 8 || password.includes(' ')) {
            loginError.textContent = "Password must be at least 8 characters long and cannot contain spaces.";
            loginError.style.display = 'block';
            return;
        }
        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        })
            .then(response => {
                if (response.ok) {
                    window.location.href = '/'; // Redirect to home page
                    return
                }
                return response.json()

            })
            .then(data => {
                console.log("Login response:", data);

                loginError.textContent = data.reason || "An error occurred during login.";
                loginError.style.display = 'block';

            })
            .catch(error => {
                console.log("Login error:", error);
                loginError.textContent = "An error occurred: " + error.message;
                loginError.style.display = 'block';
            });
    });
    registerField.addEventListener('submit', (e) => {
        e.preventDefault();
        const username = registerField.querySelector('input[name="username"]').value;
        const email = registerField.querySelector('input[name="email"]').value;
        const password = registerField.querySelector('input[name="password"]').value;
        const confirmPassword = registerField.querySelector('input[name="confirmPassword"]').value;
        registerError.style.display = 'none';
        if (username.includes(' ')) {
            registerError.textContent = "Username cannot contain spaces.";
            registerError.style.display = 'block';
            return;
        }
        if (password.length < 8 || password.includes(' ')) {
            registerError.textContent = "Password must be at least 8 characters long and cannot contain spaces.";
            registerError.style.display = 'block';
            return;
        }
        if (password !== confirmPassword) {
            registerError.textContent = "Passwords do not match.";
            registerError.style.display = 'block';
            return;
        }
        const requestData = { username, email, password, confirmPassword };
        console.log("Signup request data:", requestData);
        fetch('/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData),
        })
            .then(response => {
                if (response.ok) {
                    window.location.href = '/login'; // Redirect to home page
                    alert("please log in")
                    return
                }
                return response.json()
            })
            .then(data => {
                console.log("Signup response:", data);
                registerError.textContent = data.reason || "An error occurred during registration.";
                registerError.style.display = 'block';
            })
            .catch(error => {
                console.log("Signup error:", error);
                registerError.textContent = "An error occurred: " + error;
                registerError.style.display = 'block';
            });
    });
});
