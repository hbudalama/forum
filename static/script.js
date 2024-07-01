document.addEventListener('DOMContentLoaded', () => {
    const loginBtn = document.getElementById('loginBtn');
    const registerBtn = document.getElementById('registerBtn');
    const loginField = document.getElementById('loginField');
    const registerField = document.getElementById('registerField');

    loginBtn.addEventListener('change', () => {
        loginField.style.display = 'flex';
        registerField.style.display = 'none';
    });

    registerBtn.addEventListener('change', () => {
        loginField.style.display = 'none';
        registerField.style.display = 'flex';
    });
});
