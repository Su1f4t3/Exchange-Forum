<template>
  <div class="article-detail-container">
    <div class="container">
      <div class="article-content">
        <article v-if="article" class="article">
          <header class="article-header">
            <div class="article-meta">
              <span class="article-date">{{ formatDate(article.CreatedAt) }}</span>
              <el-tag v-if="article.category" size="small" class="article-category">
                {{ article.category }}
              </el-tag>
            </div>
            <h1 class="article-title">{{ article.title }}</h1>
            <div class="article-author-info">
              <el-avatar :size="32" :src="article.author?.avatar" class="author-avatar">
                {{ article.author?.username?.charAt(0) }}
              </el-avatar>
              <div class="author-details">
                <span class="author-name">{{ article.author?.username || '匿名用户' }}</span>
                <span class="author-bio">{{ article.author?.bio || '金融分析师' }}</span>
              </div>
            </div>
          </header>

          <div class="article-body">
            <div class="article-content markdown-body" v-html="formattedContent"></div>
          </div>

          <footer class="article-footer">
            <div class="article-stats">
              <div class="stat-item">
                <el-icon size="16"><View /></el-icon>
                <span>{{ article.views || 0 }} 阅读</span>
              </div>
              <div class="stat-item">
                <el-button
                  type="primary"
                  @click="likeArticle"
                  :loading="liking"
                  class="like-btn"
                >
                  <el-icon size="16"><StarFilled /></el-icon>
                  <span>点赞 ({{ likes }})</span>
                </el-button>
              </div>
            </div>

            <div class="article-tags" v-if="article.tags">
              <el-tag
                v-for="tag in article.tags.split(',')"
                :key="tag"
                size="small"
                class="tag-item"
              >
                {{ tag.trim() }}
              </el-tag>
            </div>
          </footer>
        </article>

        <div v-else class="auth-prompt">
          <div class="blankslate">
            <el-icon size="48" class="text-muted"><Lock /></el-icon>
            <h3>需要登录</h3>
            <p>登录后可以阅读完整的文章内容</p>
            <router-link to="/login" class="btn btn-primary">立即登录</router-link>
          </div>
        </div>
      </div>

      <aside class="article-sidebar">
        <div class="sidebar-section">
          <h3 class="sidebar-title">相关文章</h3>
          <div class="related-articles">
            <div v-for="i in 3" :key="i" class="related-article">
              <h4 class="related-title">相关文章标题 {{ i }}</h4>
              <span class="related-date">2024-01-{{ 10 + i }}</span>
            </div>
          </div>
        </div>

        <div class="sidebar-section">
          <h3 class="sidebar-title">热门标签</h3>
          <div class="tag-cloud">
            <el-tag
              v-for="tag in ['汇率', '投资', '分析', '市场', '货币']"
              :key="tag"
              size="small"
              class="cloud-tag"
            >
              {{ tag }}
            </el-tag>
          </div>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useRoute } from "vue-router";
import axios from "../axios";
import type { Article, Like } from "../types/Article";
import { Lock, View, StarFilled } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const article = ref<Article | null>(null);
const route = useRoute();
const likes = ref<number>(0);
const liking = ref(false);

const { id } = route.params;

const formattedContent = computed(() => {
  if (!article.value?.content) return '';
  // Simple markdown-like formatting
  return article.value.content
    .replace(/\n\n/g, '</p><p>')
    .replace(/\n/g, '<br>')
    .replace(/^/, '<p>')
    .replace(/$/, '</p>');
});

const fetchArticle = async () => {
  try {
    const response = await axios.get<Article>(`/articles/${id}`);
    article.value = response.data;
  } catch (error) {
    console.error("Failed to load article:", error);
    ElMessage.error('获取文章失败');
  }
};

const likeArticle = async () => {
  try {
    liking.value = true;
    const res = await axios.post<Like>(`/articles/${id}/like`);
    likes.value = res.data.likes;
    ElMessage.success('点赞成功');
  } catch (error) {
    ElMessage.error('点赞失败，请重试');
    console.log('Error Liking article:', error);
  } finally {
    liking.value = false;
  }
};

const fetchLike = async () => {
  try {
    const res = await axios.get<Like>(`/articles/${id}/like`);
    likes.value = res.data.likes;
  } catch (error) {
    console.log('Error fetching likes:', error);
  }
};

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
};

onMounted(() => {
  fetchArticle();
  fetchLike();
});
</script>

<style scoped>
.article-detail-container {
  min-height: calc(100vh - 64px);
  background-color: var(--color-canvas-default);
  padding: 48px 0;
}

.container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 24px;
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 48px;
  align-items: start;
}

.article-content {
  min-width: 0;
}

.article {
  background-color: var(--color-canvas-overlay);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-large);
  overflow: hidden;
}

.article-header {
  padding: 32px;
  background-color: var(--color-canvas-subtle);
  border-bottom: 1px solid var(--color-border-default);
}

.article-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.article-date {
  font-size: 14px;
  color: var(--color-fg-muted);
}

.article-category {
  font-size: 12px;
  font-weight: 500;
}

.article-title {
  margin: 0 0 20px;
  font-size: 32px;
  font-weight: 600;
  line-height: 1.25;
  color: var(--color-fg-default);
}

.article-author-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.author-avatar {
  border: 1px solid var(--color-border-default);
}

.author-details {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.author-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-fg-default);
}

.author-bio {
  font-size: 12px;
  color: var(--color-fg-muted);
}

.article-body {
  padding: 32px;
}

.article-content {
  font-size: 16px;
  line-height: 1.75;
  color: var(--color-fg-default);
}

.article-content p {
  margin-bottom: 16px;
}

.article-footer {
  padding: 24px 32px;
  background-color: var(--color-canvas-subtle);
  border-top: 1px solid var(--color-border-default);
}

.article-stats {
  display: flex;
  align-items: center;
  gap: 24px;
  margin-bottom: 16px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--color-fg-muted);
}

.like-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  font-size: 14px;
  font-weight: 500;
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag-item {
  font-size: 12px;
  font-weight: 500;
}

.auth-prompt {
  max-width: 600px;
  margin: 0 auto;
}

.blankslate {
  padding: 80px 40px;
  text-align: center;
  background-color: var(--color-canvas-subtle);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-large);
}

.blankslate h3 {
  margin: 16px 0 8px;
  font-size: 24px;
  font-weight: 600;
}

.blankslate p {
  margin-bottom: 24px;
  color: var(--color-fg-muted);
}

.article-sidebar {
  position: sticky;
  top: 24px;
}

.sidebar-section {
  background-color: var(--color-canvas-overlay);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-large);
  padding: 20px;
  margin-bottom: 24px;
}

.sidebar-title {
  margin: 0 0 16px;
  font-size: 16px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.related-articles {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.related-article {
  padding: 12px 0;
  border-bottom: 1px solid var(--color-border-muted);
}

.related-article:last-child {
  border-bottom: none;
}

.related-title {
  margin: 0 0 4px;
  font-size: 14px;
  font-weight: 500;
  color: var(--color-accent-fg);
  cursor: pointer;
}

.related-title:hover {
  text-decoration: underline;
}

.related-date {
  font-size: 12px;
  color: var(--color-fg-muted);
}

.tag-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.cloud-tag {
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease-out;
}

.cloud-tag:hover {
  transform: translateY(-1px);
}

@media (max-width: 1024px) {
  .container {
    grid-template-columns: 1fr;
  }

  .article-sidebar {
    position: static;
    order: -1;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 24px;
  }
}

@media (max-width: 768px) {
  .article-sidebar {
    grid-template-columns: 1fr;
  }

  .article-header,
  .article-body,
  .article-footer {
    padding: 24px;
  }

  .article-title {
    font-size: 24px;
  }
}
</style>
