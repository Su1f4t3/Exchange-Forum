<template>
  <div class="exchange-container">
    <div class="container">
      <div class="exchange-header">
        <h1 class="page-title">货币兑换</h1>
        <p class="page-description">实时汇率换算，支持全球主要货币</p>
      </div>

      <div class="exchange-card">
        <el-form :model="form" class="exchange-form" label-position="top">
          <div class="form-row">
            <div class="form-col">
              <el-form-item label="从货币" class="form-item">
                <el-select
                  v-model="form.fromCurrency"
                  placeholder="选择源货币"
                  class="currency-select"
                  size="large"
                >
                  <el-option
                    v-for="currency in currencies"
                    :key="currency"
                    :label="currency"
                    :value="currency"
                  />
                </el-select>
              </el-form-item>
            </div>

            <div class="exchange-arrow">
              <el-button circle @click="swapCurrencies" class="swap-btn">
                <el-icon size="20"><Sort /></el-icon>
              </el-button>
            </div>

            <div class="form-col">
              <el-form-item label="到货币" class="form-item">
                <el-select
                  v-model="form.toCurrency"
                  placeholder="选择目标货币"
                  class="currency-select"
                  size="large"
                >
                  <el-option
                    v-for="currency in currencies"
                    :key="currency"
                    :label="currency"
                    :value="currency"
                  />
                </el-select>
              </el-form-item>
            </div>
          </div>

          <el-form-item label="兑换金额" class="form-item">
            <el-input
              v-model="form.amount"
              type="number"
              placeholder="输入要兑换的金额"
              class="amount-input"
              size="large"
            >
              <template #prefix>
                <span class="currency-prefix">{{ form.fromCurrency || 'USD' }}</span>
              </template>
            </el-input>
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              @click="exchange"
              class="exchange-btn"
              size="large"
              :loading="loading"
              :disabled="!form.fromCurrency || !form.toCurrency || !form.amount"
            >
              立即兑换
            </el-button>
          </el-form-item>
        </el-form>

        <div v-if="result" class="result-box">
          <div class="result-header">
            <h3 class="result-title">兑换结果</h3>
            <span class="result-rate">汇率: {{ displayRate }}</span>
          </div>
          <div class="result-content">
            <div class="result-amount">
              <span class="amount-from">{{ formatAmount(form.amount) }} {{ form.fromCurrency }}</span>
              <el-icon class="arrow-icon"><Right /></el-icon>
              <span class="amount-to">{{ formatAmount(result) }} {{ form.toCurrency }}</span>
            </div>
            <div class="result-time">
              <el-icon size="14"><Clock /></el-icon>
              更新时间: {{ currentTime }}
            </div>
          </div>
        </div>
      </div>

      <div v-if="recentRates.length > 0" class="recent-rates">
        <h2 class="section-title">最近汇率</h2>
        <div class="rates-grid">
          <div v-for="rate in recentRates" :key="rate.key" class="rate-item">
            <div class="rate-currency">{{ rate.from }} → {{ rate.to }}</div>
            <div class="rate-value">{{ rate.rate.toFixed(4) }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>  
  
<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { Sort, Right, Clock } from '@element-plus/icons-vue';
import axios from '../axios';
import { ElMessage } from 'element-plus';

interface ExchangeRate {
  fromcurrency: string;
  tocurrency: string;
  rate: number;
}

const form = ref({
  fromCurrency: 'USD',
  toCurrency: 'CNY',
  amount: 100,
});

const result = ref<number | null>(null);
const currencies = ref<string[]>([]);
const rates = ref<ExchangeRate[]>([]);
const loading = ref(false);
const currentRate = ref<number | null>(null);
const currentTime = ref('');

const recentRates = computed(() => {
  if (!rates.value.length) return [];
  return rates.value.slice(0, 6).map(rate => ({
    key: `${rate.fromcurrency}-${rate.tocurrency}`,
    from: rate.fromcurrency,
    to: rate.tocurrency,
    rate: rate.rate
  }));
});

const displayRate = computed(() => {
  return currentRate.value ? currentRate.value.toFixed(4) : '-';
});

const fetchCurrencies = async () => {
  try {
    loading.value = true;
    const response = await axios.get<ExchangeRate[]>('/exchangerates');
    rates.value = response.data;
    currencies.value = [...new Set(response.data.map((rate: ExchangeRate) => [rate.fromcurrency, rate.tocurrency]).flat())].sort();
  } catch (error) {
    ElMessage.error('获取汇率数据失败');
    console.error('Failed to load currencies', error);
  } finally {
    loading.value = false;
  }
};

const exchange = () => {
  if (!form.value.amount || form.value.amount <= 0) {
    ElMessage.warning('请输入有效的兑换金额');
    return;
  }

  const rate = rates.value.find(
    (rate) => rate.fromcurrency === form.value.fromCurrency && rate.tocurrency === form.value.toCurrency
  )?.rate;

  if (rate) {
    currentRate.value = rate;
    result.value = form.value.amount * rate;
    updateCurrentTime();
  } else {
    ElMessage.warning('未找到对应的汇率信息');
    result.value = null;
    currentRate.value = null;
  }
};

const swapCurrencies = () => {
  const temp = form.value.fromCurrency;
  form.value.fromCurrency = form.value.toCurrency;
  form.value.toCurrency = temp;

  if (result.value) {
    exchange();
  }
};

const formatAmount = (amount: number) => {
  return new Intl.NumberFormat('zh-CN', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(amount);
};

const updateCurrentTime = () => {
  currentTime.value = new Date().toLocaleString('zh-CN');
};

onMounted(() => {
  fetchCurrencies();
  updateCurrentTime();
});
</script>
  
<style scoped>
.exchange-container {
  min-height: calc(100vh - 64px);
  background-color: var(--color-canvas-default);
}

.container {
  max-width: 768px;
  margin: 0 auto;
  padding: 48px 24px;
}

.exchange-header {
  text-align: center;
  margin-bottom: 48px;
}

.page-title {
  margin-bottom: 8px;
  font-size: 32px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.page-description {
  font-size: 18px;
  color: var(--color-fg-muted);
}

.exchange-card {
  background-color: var(--color-canvas-overlay);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-medium);
  padding: 32px;
  margin-bottom: 32px;
}

.exchange-form {
  margin-bottom: 32px;
}

.form-row {
  display: flex;
  align-items: flex-end;
  gap: 16px;
  margin-bottom: 24px;
}

.form-col {
  flex: 1;
}

.exchange-arrow {
  padding-bottom: 24px;
}

.swap-btn {
  width: 40px;
  height: 40px;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-canvas-subtle);
  border: 1px solid var(--color-border-default);
  color: var(--color-fg-muted);
  transition: all 0.2s ease-out;
}

.swap-btn:hover {
  background-color: var(--color-canvas-default);
  border-color: var(--color-neutral-emphasis);
  color: var(--color-fg-default);
}

.form-item {
  margin-bottom: 0;
}

.form-item :deep(.el-form-item__label) {
  color: var(--color-fg-default);
  font-weight: 500;
  margin-bottom: 8px;
  font-size: 14px;
}

.currency-select {
  width: 100%;
}

.currency-prefix {
  color: var(--color-fg-muted);
  font-size: 14px;
  margin-right: 4px;
}

.amount-input {
  width: 100%;
}

.exchange-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
  margin-top: 8px;
}

.result-box {
  background-color: var(--color-canvas-subtle);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-medium);
  padding: 24px;
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.result-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.result-rate {
  font-size: 14px;
  color: var(--color-fg-muted);
}

.result-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.result-amount {
  display: flex;
  align-items: center;
  gap: 16px;
  font-size: 20px;
  font-weight: 600;
}

.amount-from {
  color: var(--color-fg-muted);
}

.amount-to {
  color: var(--color-success-fg);
}

.arrow-icon {
  color: var(--color-fg-subtle);
}

.result-time {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: var(--color-fg-muted);
}

.recent-rates {
  margin-top: 48px;
}

.section-title {
  margin-bottom: 24px;
  font-size: 24px;
  font-weight: 600;
  color: var(--color-fg-default);
}

.rates-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

.rate-item {
  padding: 16px;
  background-color: var(--color-canvas-overlay);
  border: 1px solid var(--color-border-default);
  border-radius: var(--border-radius-medium);
  transition: all 0.2s ease-out;
}

.rate-item:hover {
  border-color: var(--color-accent-muted);
  box-shadow: var(--shadow-small);
}

.rate-currency {
  font-size: 14px;
  color: var(--color-fg-muted);
  margin-bottom: 8px;
}

.rate-value {
  font-size: 20px;
  font-weight: 600;
  color: var(--color-fg-default);
}

@media (max-width: 768px) {
  .container {
    padding: 32px 16px;
  }

  .exchange-card {
    padding: 24px;
  }

  .form-row {
    flex-direction: column;
    align-items: stretch;
  }

  .exchange-arrow {
    display: flex;
    justify-content: center;
    padding: 8px 0;
    transform: rotate(90deg);
  }

  .result-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .result-amount {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .arrow-icon {
    transform: rotate(90deg);
  }
}
</style>
  