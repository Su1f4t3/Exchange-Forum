<template>
  <div class="auth-page">
    <div class="auth-container">
      <div class="auth-card">
        <div class="auth-header">
          <h1 class="auth-title">注册新账户</h1>
          <p class="auth-subtitle">加入 Exchange Forum，获取专业金融服务</p>
        </div>

        <el-form :model="form" class="auth-form" @submit.prevent="register" label-position="top">
          <el-form-item label="用户名" class="form-item" required>
            <el-input
              v-model="form.username"
              placeholder="请输入用户名"
              size="large"
              class="auth-input"
              prefix-icon="User"
            />
          </el-form-item>

          <el-form-item label="密码" class="form-item" required>
            <el-input
              v-model="form.password"
              type="password"
              placeholder="请输入密码（至少6位）"
              size="large"
              class="auth-input"
              prefix-icon="Lock"
              show-password
            />
          </el-form-item>

          <el-form-item label="确认密码" class="form-item" required>
            <el-input
              v-model="form.confirmPassword"
              type="password"
              placeholder="请再次输入密码"
              size="large"
              class="auth-input"
              prefix-icon="Lock"
              show-password
            />
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              native-type="submit"
              class="auth-btn"
              size="large"
              :loading="loading"
              :disabled="!isFormValid"
            >
              创建账户
            </el-button>
          </el-form-item>
        </el-form>

        <div class="auth-footer">
          <p class="auth-text">
            已有账户？
            <router-link to="/login" class="auth-link">立即登录</router-link>
          </p>
        </div>
      </div>

      <div class="auth-info">
        <h2 class="info-title">加入 Exchange Forum</h2>
        <p class="info-description">创建免费账户，解锁全部功能</p>
        <ul class="info-features">
          <li>实时汇率查询</li>
          <li>货币兑换计算</li>
          <li>金融新闻资讯</li>
          <li>个人收藏管理</li>
        </ul>
      </div>
    </div>
  </div>
</template>
  
<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';
import { ElMessage } from 'element-plus';

const form = ref({
  username: '',
  password: '',
  confirmPassword: '',
});

const loading = ref(false);
const authStore = useAuthStore();
const router = useRouter();

const isFormValid = computed(() => {
  return form.value.username &&
         form.value.password &&
         form.value.confirmPassword &&
         form.value.password === form.value.confirmPassword &&
         form.value.password.length >= 6;
});

const register = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    ElMessage.error('两次输入的密码不一致');
    return;
  }

  if (form.value.password.length < 6) {
    ElMessage.error('密码长度至少为6位');
    return;
  }

  try {
    loading.value = true;
    await authStore.register(form.value.username, form.value.password);
    ElMessage.success('注册成功');
    router.push({ name: 'News' });
  } catch (error) {
    ElMessage.error('注册失败，请重试');
  } finally {
    loading.value = false;
  }
};
</script>
  
<style scoped>
.auth-page {
  min-height: 100vh;
  background-color: var(--color-canvas-default);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 48px 24px;
}

.auth-container {
  display: flex;
  gap: 80px;
  max-width: 1024px;
  width: 100%;
}

.auth-card {
  flex: 1;
  max-width: 400px;
  background-color: var(--color-canvas-overlay);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-medium);
  padding: 32px;
}

.auth-header {
  text-align: center;
  margin-bottom: 32px;
}


.auth-title {
  margin: 0 0 8px;
  font-size: 24px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.auth-subtitle {
  margin: 0;
  font-size: 16px;
  color: var(--color-fg-muted);
}

.auth-form {
  margin-bottom: 24px;
}

.form-item {
  margin-bottom: 20px;
}

.form-item :deep(.el-form-item__label) {
  color: var(--color-fg-default);
  font-weight: 500;
  margin-bottom: 8px;
  font-size: 14px;
}

.auth-input :deep(.el-input__wrapper) {
  width: 100%;
}

.auth-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
}

.auth-footer {
  text-align: center;
  padding-top: 24px;
  border-top: 1px solid var(--color-border-muted);
}

.auth-text {
  margin: 0;
  font-size: 14px;
  color: var(--color-fg-muted);
}

.auth-link {
  color: var(--color-accent-fg);
  text-decoration: none;
  font-weight: 500;
}

.auth-link:hover {
  text-decoration: underline;
}

.auth-info {
  flex: 1;
  max-width: 400px;
  padding: 32px 0;
}

.info-title {
  margin: 0 0 16px;
  font-size: 32px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.info-description {
  margin: 0 0 32px;
  font-size: 18px;
  line-height: 1.5;
  color: var(--color-fg-muted);
}

.info-features {
  list-style: none;
  padding: 0;
  margin: 0;
}

.info-features li {
  position: relative;
  padding-left: 28px;
  margin-bottom: 12px;
  font-size: 16px;
  color: var(--color-fg-default);
}

.info-features li::before {
  content: "✓";
  position: absolute;
  left: 0;
  top: 0;
  color: var(--color-success-fg);
  font-weight: 600;
}

@media (max-width: 768px) {
  .auth-container {
    flex-direction: column;
    gap: 48px;
  }

  .auth-info {
    text-align: center;
    padding: 0;
  }

  .info-features li {
    text-align: left;
  }
}
</style>
  