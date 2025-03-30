<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const form = ref({
  username: '',
  phone: '',
  password: '',
  confirmPassword: ''
})
const error = ref('')
const isLoading = ref(false)
const router = useRouter()

const submitForm = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    error.value = '两次输入的密码不一致'
    return
  }
  
  isLoading.value = true
  try {
    const response = await axios.post('http://127.0.0.1:9001/api/register', {
      username: form.value.username,
      phone: form.value.phone,
      password: form.value.password
    })
    
    // 注册成功后自动登录
    localStorage.setItem('token', response.data.token)
    document.querySelector('.register-container').classList.add('success')
    setTimeout(() => {
      router.push('/dashboard')
    }, 1000)
  } catch (err) {
    error.value = err.response?.data?.error || '注册失败'
    document.querySelector('.register-container').classList.add('shake')
    setTimeout(() => {
      document.querySelector('.register-container').classList.remove('shake')
    }, 500)
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="register-container">
    <div class="register-card">
      <div class="register-header">
        <h2>创建新账户</h2>
        <p>填写以下信息完成注册</p>
      </div>
      
      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="username">用户名</label>
          <input v-model="form.username" type="text" id="username" placeholder="请输入用户名" required>
        </div>
        
        <div class="form-group">
          <label for="phone">手机号</label>
          <input v-model="form.phone" type="tel" id="phone" placeholder="请输入手机号" required>
        </div>

        <div class="form-group">
          <label for="password">密码</label>
          <input v-model="form.password" type="password" id="password" placeholder="请输入密码" required>
        </div>

        <div class="form-group">
          <label for="confirmPassword">确认密码</label>
          <input v-model="form.confirmPassword" type="password" id="confirmPassword" placeholder="请再次输入密码" required>
        </div>
        
        <button type="submit" :disabled="isLoading">
          <span v-if="!isLoading">注 册</span>
          <span v-else class="loader"></span>
        </button>
        
        <p v-if="error" class="error">{{ error }}</p>
      </form>
      
      <div class="register-footer">
        <p>已有账号? <router-link to="/login">立即登录</router-link></p>
      </div>
    </div>
    
    <div class="particles"></div>
  </div>
</template>

<style scoped>
.register-container {
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

.register-container.success {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.register-card {
  width: 90%;
  max-width: 500px;
  min-width: 300px;
  padding: 2.5rem;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(8px);
  z-index: 1;
  transform: translateY(0);
  transition: all 0.3s ease;
  margin: 1rem;
}

@media (min-width: 768px) {
  .register-card {
    padding: 3rem;
    width: 80%;
  }
}

@media (min-width: 1024px) {
  .register-card {
    width: 50%;
  }
}

.register-header {
  text-align: center;
  margin-bottom: 30px;
}

.register-header h2 {
  color: #333;
  font-size: 28px;
  margin-bottom: 8px;
}

.register-header p {
  color: #666;
  font-size: 14px;
}

.form-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
}

.form-group label {
  margin-bottom: 5px;
  font-size: 14px;
  color: #666;
}

.form-group input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: #f9f9f9;
  font-size: 1rem;
  color: #333;
  transition: all 0.3s ease;
}

@media (min-width: 768px) {
  .form-group input {
    padding: 1rem;
  }
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
}

button:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

button:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.error {
  color: #ff4757;
  text-align: center;
  margin-top: 15px;
  font-size: 14px;
}

.register-footer {
  text-align: center;
  margin-top: 20px;
  font-size: 14px;
  color: #666;
}

.register-footer a {
  color: #667eea;
  text-decoration: none;
}
</style>