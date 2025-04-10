<template>
  <div class="scan-container">
    <form @submit.prevent="submitUrl" class="scan-form">
      <h2>网站扫描</h2>
      <div class="form-group">
        <label for="targetUrl">请输入目标URL:</label>
        <input
          type="url"
          id="targetUrl"
          v-model="targetUrl"
          placeholder="https://example.com"
          required
          class="url-input"
        />
      </div>
      <button type="submit" class="submit-btn">开始扫描</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const targetUrl = ref('')

const submitUrl = async () => {
  try {
    const response = await axios.post('/api/scan', {
      url: targetUrl.value
    })
    console.log('扫描结果:', response.data)
    // 这里可以添加处理扫描结果的逻辑
  } catch (error) {
    console.error('扫描失败:', error)
    // 这里可以添加错误处理逻辑
  }
}
</script>

<style scoped>
.scan-container {
  display: flex;
  justify-content: center;
  padding-top: 2rem;
}

.scan-form {
  width: 100%;
  max-width: 600px;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
}

.url-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ced4da;
  border-radius: 4px;
  font-size: 1rem;
}

.submit-btn {
  background-color: #007bff;
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.2s;
}

.submit-btn:hover {
  background-color: #0056b3;
}
</style>