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

        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert("Login successful!");
            } else {
                loginError.textContent = data.message;
                loginError.style.display = 'block';
            }
        })
        .catch(error => {
            loginError.textContent = "An error occurred. Please try again.";
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

        if (password !== confirmPassword) {
            registerError.textContent = "Passwords do not match.";
            registerError.style.display = 'block';
            return;
        }

        fetch('/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, email, password, confirmPassword }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert("Registration successful!");
            } else {
                registerError.textContent = data.message;
                registerError.style.display = 'block';
            }
        })
        .catch(error => {
            registerError.textContent = "An error occurred. Please try again.";
            registerError.style.display = 'block';
        });
    });
});
