<template>
  <div class="modal-overlay" @click="close">
    <div class="modal">
      <h2>编辑个人资料</h2>
      
      <div class="avatar-upload">
        <img :src="tempAvatar || user.avatar || '/default-avatar.png'" class="preview" />
        <input 
          type="file" 
          accept="image/*"
          @change="handleAvatarChange"
          ref="fileInput"
        />
        <button @click="$refs.fileInput.click()">选择头像</button>
      </div>

      <form @submit.prevent="submit">
        <div class="form-group">
          <label>用户名</label>
          <input v-model="form.username" />
        </div>
        
        <div class="form-group">
          <label>手机号</label>
          <input v-model="form.phone" />
        </div>

        <div class="actions">
          <button type="button" @click="close">取消</button>
          <button type="submit">保存</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'

export default {
  props: {
    user: {
      type: Object,
      required: true
    }
  },

  emits: ['close', 'update'],

  setup(props, { emit }) {
    const authStore = useAuthStore()
    const form = ref({ ...props.user })
    const tempAvatar = ref(null)

    const handleAvatarChange = (e) => {
      const file = e.target.files[0]
      if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
          tempAvatar.value = e.target.result
        }
        reader.readAsDataURL(file)
      }
    }

    const submit = async () => {
      try {
        const updatedUser = await authStore.updateProfile({
          ...form.value,
          avatar: tempAvatar.value
        })
        emit('update', updatedUser)
        close()
      } catch (error) {
        console.error('更新失败:', error)
      }
    }

    const close = () => {
      emit('close')
    }

    return {
      form,
      tempAvatar,
      handleAvatarChange,
      submit,
      close
    }
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 400px;
  max-width: 90%;
}

.avatar-upload {
  margin-bottom: 1rem;
  text-align: center;
}

.preview {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 0.5rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
}

.form-group input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
  margin-top: 1rem;
}

.actions button {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.actions button[type="button"] {
  background: #f5f5f5;
}

.actions button[type="submit"] {
  background: #4CAF50;
  color: white;
}
</style>