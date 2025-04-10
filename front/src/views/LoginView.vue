<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const form = ref({
  username: '',
  password: ''
})
const error = ref('')
const isLoading = ref(false)
const router = useRouter()

const submitForm = async () => {
  isLoading.value = true
  try {
    const response = await axios.post('http://127.0.0.1:9001/api/login', form.value)
    localStorage.setItem('token', response.data.token)
    const authStore = useAuthStore()
    authStore.login(response.data.user)

    document.querySelector('.login-container').classList.add('success')
    setTimeout(() => {
      router.push('/dashboard')
    }, 1000)
  } catch (err) {
    error.value = err.response?.data?.error || '登录失败'
    document.querySelector('.login-container').classList.add('shake')
    setTimeout(() => {
      document.querySelector('.login-container').classList.remove('shake')
    }, 500)
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h2>欢迎回来</h2>
        <p>请登录您的账户</p>
      </div>

      <form @submit.prevent="submitForm">
        <div class="form-group floating">
          <input v-model="form.username" type="text" id="username" required>
          <label for="username">用户名</label>
          <div class="underline"></div>
        </div>

        <div class="form-group floating">
          <input v-model="form.password" type="password" id="password" required>
          <label for="password">密码</label>
          <div class="underline"></div>
        </div>

        <button type="submit" :disabled="isLoading">
          <span v-if="!isLoading">登 录</span>
          <span v-else class="loader"></span>
        </button>

        <p v-if="error" class="error">{{ error }}</p>
      </form>

    <div class="login-footer">
        <p>还没有账号? <router-link to="/register">立即注册</router-link></p>
    </div>
    </div>

    <div class="particles"></div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
  overflow: hidden;
  transition: all 0.5s ease;
  margin: 0;
  padding: 0;
}

.particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><circle cx="25" cy="25" r="1" fill="rgba(255,255,255,0.3)"/><circle cx="75" cy="25" r="1.5" fill="rgba(255,255,255,0.3)"/><circle cx="50" cy="50" r="1" fill="rgba(255,255,255,0.3)"/><circle cx="25" cy="75" r="1.5" fill="rgba(255,255,255,0.3)"/><circle cx="75" cy="75" r="1" fill="rgba(255,255,255,0.3)"/></svg>');
  background-size: 50px 50px;
  animation: particlesMove 20s linear infinite;
  z-index: 0;
}

.login-container.success {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.login-card {
  width: 90%;
  max-width: 500px;
  min-width: 300px;
  padding: 2.5rem;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(8px);
  z-index: 1;
  transform: translateY(0);
  transition: all 0.3s ease;
  margin: 1rem;
}

@media (min-width: 768px) {
  .login-card {
    padding: 3rem;
    width: 80%;
  }
}

@media (min-width: 1024px) {
  .login-card {
    width: 50%;
  }
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-header h2 {
  color: #333;
  font-size: 28px;
  margin-bottom: 8px;
}

.login-header p {
  color: #666;
  font-size: 14px;
}

.form-group {
  position: relative;
  margin-bottom: 25px;
}

.form-group.floating label {
  position: absolute;
  top: 0;
  left: 0;
  color: #999;
  pointer-events: none;
  transform-origin: 0 0;
  transition: all 0.2s ease;
}

.form-group.floating input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: #fff;
  font-size: 1rem;
  color: #333;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  transition: all 0.3s ease;
}

@media (min-width: 768px) {
  .form-group.floating input {
    padding: 1rem 1.25rem;
  }
}

.form-group.floating input:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
  outline: none;
}

.form-group.floating input:focus + label,
.form-group.floating input:not(:placeholder-shown) + label {
  transform: translateY(-20px) scale(0.8);
  color: #667eea;
}

.underline {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 2px;
  background: #667eea;
  transition: all 0.3s ease;
}

.form-group.floating input:focus ~ .underline {
  width: 100%;
}

button {
  width: 100%;
  padding: 12px;
  background: linear-gradient(to right, #667eea, #764ba2);
  color: white;
  border: none;
  border-radius: 25px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

button:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

button:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.loader {
  display: inline-block;
  width: 20px;
  height: 20px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: white;
  animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.error {
  color: #ff4757;
  text-align: center;
  margin-top: 15px;
  font-size: 14px;
}

.login-footer {
  text-align: center;
  margin-top: 20px;
  font-size: 14px;
  color: #666;
}

.login-footer a {
  color: #667eea;
  text-decoration: none;
}

.particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><circle cx="25" cy="25" r="1" fill="rgba(255,255,255,0.3)"/><circle cx="75" cy="25" r="1.5" fill="rgba(255,255,255,0.3)"/><circle cx="50" cy="50" r="1" fill="rgba(255,255,255,0.3)"/><circle cx="25" cy="75" r="1.5" fill="rgba(255,255,255,0.3)"/><circle cx="75" cy="75" r="1" fill="rgba(255,255,255,0.3)"/></svg>');
  background-size: 50px 50px;
  animation: particlesMove 20s linear infinite;
}

@keyframes particlesMove {
  from { background-position: 0 0; }
  to { background-position: 50px 50px; }
}

/* 震动动画 */
.shake {
  animation: shake 0.5s cubic-bezier(.36,.07,.19,.97) both;
}

@keyframes shake {
  10%, 90% { transform: translateX(-1px); }
  20%, 80% { transform: translateX(2px); }
  30%, 50%, 70% { transform: translateX(-4px); }
  40%, 60% { transform: translateX(4px); }
}
</style>