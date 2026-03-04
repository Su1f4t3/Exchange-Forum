<template>
  <div class="home-container">
    <div class="hero-section">
      <div class="hero-content">
        <h1 class="hero-title">欢迎使用 Exchange Forum</h1>
        <p class="hero-description">实时汇率查询、货币兑换计算与金融新闻资讯，一站式金融服务平台</p>
        <div class="hero-actions">
          <router-link to="/exchanges" class="btn btn-primary hero-btn">
            开始汇率兑换
          </router-link>
          <router-link to="/articles" class="btn btn-secondary hero-btn">
            浏览金融新闻
          </router-link>
        </div>
      </div>
    </div>

    <div class="features-section">
      <div class="container">
        <h2 class="section-title">核心功能</h2>
        <div class="features-grid">
          <div class="feature-card">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 16 16" fill="currentColor">
                <path d="M1.5 1.75a.75.75 0 00-1.5 0v12.5c0 .414.336.75.75.75h14.5a.75.75 0 000-1.5H1.5V1.75zm2.75 7a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5z"></path>
              </svg>
            </div>
            <h3 class="feature-title">实时汇率</h3>
            <p class="feature-description">获取全球主要货币的实时汇率数据，支持历史汇率查询</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 16 16" fill="currentColor">
                <path d="M8.5.75a.75.75 0 00-1.5 0v5.19L4.391 3.33a.75.75 0 10-1.06 1.061L5.939 7H.75a.75.75 0 000 1.5h5.19l-2.61 2.609a.75.75 0 101.061 1.06L7 9.561v5.189a.75.75 0 001.5 0V9.56l2.609 2.61a.75.75 0 101.06-1.061L9.561 8.5h5.189a.75.75 0 000-1.5H9.56l2.61-2.609a.75.75 0 00-1.061-1.06L8.5 5.939V.75z"></path>
              </svg>
            </div>
            <h3 class="feature-title">货币兑换</h3>
            <p class="feature-description">快速准确的货币兑换计算器，支持多种货币实时换算</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              <svg width="24" height="24" viewBox="0 0 16 16" fill="currentColor">
                <path d="M0 1.75A.75.75 0 01.75 1h4.253c1.227 0 2.317.59 3 1.501A3.744 3.744 0 0111.006 1h4.245a.75.75 0 01.75.75v10.5a.75.75 0 01-.75.75h-4.507a2.25 2.25 0 00-1.591.659l-.622.621a.75.75 0 01-1.06 0l-.622-.621A2.25 2.25 0 005.258 13H.75a.75.75 0 01-.75-.75V1.75zm8.755 3a2.25 2.25 0 012.25-2.25H14.5v9.25h-3.757c-.71 0-1.4.201-1.992.572l.004-7.322zm-1.504 7.324l.004-5.073-.002-2.253A2.25 2.25 0 005.003 2.5H1.5v9.25h3.757a3.75 3.75 0 011.994.574z"></path>
              </svg>
            </div>
            <h3 class="feature-title">金融资讯</h3>
            <p class="feature-description">最新的金融新闻和市场分析，助您把握投资机会</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.home-container {
  min-height: calc(100vh - 64px);
  background-color: var(--color-canvas-default);
}

.hero-section {
  padding: 96px 0;
  text-align: center;
  background: linear-gradient(135deg, var(--color-canvas-subtle) 0%, var(--color-canvas-default) 100%);
  border-bottom: 1px solid var(--color-border-default);
}

.hero-content {
  max-width: 768px;
  margin: 0 auto;
  padding: 0 24px;
}


.hero-title {
  margin-bottom: 16px;
  font-size: 48px;
  font-weight: 600;
  line-height: 1.2;
  color: var(--color-fg-default);
}

.hero-description {
  margin-bottom: 32px;
  font-size: 20px;
  line-height: 1.5;
  color: var(--color-fg-muted);
}

.hero-actions {
  display: flex;
  gap: 16px;
  justify-content: center;
  flex-wrap: wrap;
}

.hero-btn {
  padding: 12px 24px;
  font-size: 16px;
  font-weight: 500;
}

.features-section {
  padding: 80px 0;
}

.container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 24px;
}

.section-title {
  margin-bottom: 48px;
  font-size: 32px;
  font-weight: 600;
  text-align: center;
  color: var(--color-fg-default);
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 32px;
}

.feature-card {
  padding: 32px;
  background-color: var(--color-canvas-default);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-large);
  text-align: center;
  transition: all 0.2s ease-out;
}

.feature-card:hover {
  border-color: var(--color-accent-muted);
  box-shadow: var(--shadow-medium);
  transform: translateY(-2px);
}

.feature-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  margin-bottom: 16px;
  color: var(--color-accent-fg);
  background-color: var(--color-accent-subtle);
  border-radius: var(--border-radius-medium);
}

.feature-title {
  margin-bottom: 8px;
  font-size: 20px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.feature-description {
  color: var(--color-fg-muted);
  line-height: 1.5;
}

@media (max-width: 768px) {
  .hero-section {
    padding: 64px 0;
  }

  .hero-title {
    font-size: 36px;
  }

  .hero-description {
    font-size: 18px;
  }

  .hero-actions {
    flex-direction: column;
    align-items: center;
  }

  .hero-btn {
    width: 100%;
    max-width: 300px;
  }

  .features-section {
    padding: 64px 0;
  }

  .section-title {
    font-size: 28px;
  }

  .features-grid {
    grid-template-columns: 1fr;
    gap: 24px;
  }
}
</style>