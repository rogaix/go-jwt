<template>
  <div class="dashboard">
    <h1>Welcome to Dashboard</h1>
    <p>You are logged in as: {{ auth.user?.email || 'User' }}</p>
    <button @click="handleLogout" class="logout-btn">Logout</button>
  </div>
</template>

<script setup>
import { useAuthStore } from '../store/auth'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

const handleLogout = () => {
  auth.logout()
  router.push('/login')
}
</script>

<style scoped>
.dashboard {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.logout-btn {
  background: #ff4444;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
}

.logout-btn:hover {
  background: #cc0000;
}
</style> 