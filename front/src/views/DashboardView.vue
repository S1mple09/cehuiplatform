<template>
  <div class="dashboard">
    <!-- 顶部导航栏 -->
    <header class="header">
      <div class="search-bar">
        <input 
          type="text" 
          placeholder="搜索..." 
          v-model="searchQuery"
          @keyup.enter="handleSearch"
        />
        <button @click="handleSearch">搜索</button>
      </div>
      
      <div class="user-profile">
        <img 
          :src="user.avatar || '/default-avatar.png'" 
          alt="用户头像"
          class="avatar"
          @click="showProfileModal = true"
        />
        <span class="username">{{ user.username }}</span>
      </div>
    </header>

    <!-- 主要内容区域 -->
    <main class="content">
      <!-- 这里放仪表盘主要内容 -->
    </main>

    <!-- 用户信息模态框 -->
    <ProfileModal 
      v-if="showProfileModal"
      :user="user"
      @close="showProfileModal = false"
      @update="handleProfileUpdate"
    />
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import ProfileModal from '../components/ProfileModal.vue'

export default {
  components: { ProfileModal },
  setup() {
    const authStore = useAuthStore()
    const user = ref({})
    const searchQuery = ref('')
    const showProfileModal = ref(false)

    onMounted(() => {
      user.value = authStore.user
    })

    const handleSearch = () => {
      // 实现搜索逻辑
      console.log('搜索:', searchQuery.value)
    }

    const handleProfileUpdate = (updatedUser) => {
      user.value = updatedUser
    }

    return {
      user,
      searchQuery,
      showProfileModal,
      handleSearch,
      handleProfileUpdate
    }
  }
}
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  height: 100vh;
  min-width: 1200px;
  max-width: 1600px;
  margin: 0 auto;
  width: 95%;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  width: 100%;
}

.search-bar {
  display: flex;
  gap: 0.5rem;
  width: 40%;
}

.search-bar input {
  width: 100%;
  padding: 0.5rem;
}

.user-profile {
  display: flex;
  align-items: center;
  gap: 1rem;
  cursor: pointer;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
}

.username {
  font-size: 1.1rem;
  font-weight: 500;
}

.content {
  flex: 1;
  padding: 2rem;
  width: 100%;
}

@media (max-width: 1400px) {
  .dashboard {
    min-width: 1000px;
    max-width: 1200px;
  }
}

@media (max-width: 1100px) {
  .dashboard {
    min-width: 800px;
    max-width: 1000px;
  }
  
  .content {
    padding: 1.5rem;
  }
}
</style>