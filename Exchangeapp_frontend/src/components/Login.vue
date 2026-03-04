<template>
  <div class="auth-page">
    <div class="auth-container">
      <div class="auth-card">
        <div class="auth-header">
          <h1 class="auth-title">登录到 Exchange Forum</h1>
          <p class="auth-subtitle">欢迎使用汇率兑换与金融资讯平台</p>
        </div>

        <el-form :model="form" class="auth-form" @submit.prevent="login" label-position="top">
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
              placeholder="请输入密码"
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
              :disabled="!form.username || !form.password"
            >
              登录
            </el-button>
          </el-form-item>
        </el-form>

        <div class="auth-footer">
          <p class="auth-text">
            还没有账户？
            <router-link to="/register" class="auth-link">立即注册</router-link>
          </p>
        </div>
      </div>

      <div class="auth-info">
        <h2 class="info-title">Exchange Forum</h2>
        <p class="info-description">专业的汇率兑换与金融资讯平台</p>
        <ul class="info-features">
          <li>实时汇率查询</li>
          <li>货币兑换计算</li>
          <li>金融新闻资讯</li>
          <li>市场分析报告</li>
        </ul>
      </div>
    </div>
  </div>
</template>  
  
<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';
import { ElMessage } from 'element-plus';

const form = ref({
  username: '',
  password: '',
});

const loading = ref(false);
const authStore = useAuthStore();
const router = useRouter();

const login = async () => {
  if (!form.value.username || !form.value.password) {
    ElMessage.warning('请填写用户名和密码');
    return;
  }

  try {
    loading.value = true;
    await authStore.login(form.value.username, form.value.password);
    ElMessage.success('登录成功');
    router.push({ name: 'News' });
  } catch (error) {
    ElMessage.error('登录失败，请检查用户名和密码');
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
  