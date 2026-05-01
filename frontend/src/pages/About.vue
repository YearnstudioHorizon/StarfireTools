<script setup>
import logo from "../assets/images/appicon.png";
import { GetVersion } from "../../wailsjs/go/main/App";
import { ref } from "vue";

const version = ref("");
GetVersion().then((data) => {
  version.value = data;
});

function openGithub() {
  window.runtime.BrowserOpenURL("https://github.com/YearnstudioHorizon");
}

function openYearnstudio() {
  window.runtime.BrowserOpenURL("https://yearn.studio");
}

function openStarFire() {
  window.runtime.BrowserOpenURL("https://space.bilibili.com/1845800720");
}
</script>

<template>
  <div class="page-div">
    <div class="header-section">
      <h1 class="page-title">关于</h1>
      <p class="page-subtitle">关于本产品</p>
    </div>

    <div class="cards-container">
      <!-- 产品信息 -->
      <div class="about-card product-card">
        <img :src="logo" alt="logo" class="product-logo" />
        <div class="product-info">
          <h2 class="product-title">Starfire Tools 星火工具箱</h2>
          <div class="product-version">版本 {{ version }}</div>
          <div class="product-desc">一个聚合系统工具与功能管理的桌面应用</div>
        </div>
      </div>

      <!-- 开发者信息 -->
      <div class="about-card">
        <div class="card-title">开发者</div>
        <div class="dev-columns">
          <div class="dev-col">
            <span class="col-label">发行商</span>
            <span class="col-value">Starfire United Team</span>
          </div>
          <div class="dev-col">
            <span class="col-label">程序</span>
            <span class="col-value">Yearnstudio</span>
          </div>
        </div>
      </div>

      <!-- 官网与支持信息 -->
      <div class="about-card support-card">
        <div class="card-title">官网</div>
        <div class="support-content">
          <div class="support-text">
            <div class="support-title">Github</div>
            <div class="support-desc">您可以在这里查看源码</div>
          </div>
          <a href="#About" class="support-link" @click="openGithub">
            访问官网
          </a>
        </div>
        <div class="support-content">
          <div class="support-text">
            <div class="support-title">Starfire United Team</div>
            <div class="support-desc">发行商</div>
          </div>
          <a href="#About" class="support-link" @click="openStarFire">
            访问官网
          </a>
        </div>
        <div class="support-content">
          <div class="support-text">
            <div class="support-title">Yearnstudio官网</div>
            <div class="support-desc">代码团队</div>
          </div>
          <a href="#About" class="support-link" @click="openYearnstudio">
            访问官网
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-family:
    -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial,
    sans-serif;
}

.page-div {
  margin: 20px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.header-section {
  width: 95%;
  margin-bottom: 24px;
}

.page-title {
  color: var(--font-color);
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 6px;
}

.page-subtitle {
  font-size: 14px;
  color: #606060;
}

.cards-container {
  width: 95%;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 基础样式 */
.about-card {
  --win11-radius: 8px;
  --win11-border-color: rgba(0, 0, 0, 0.08);
  --win11-bg: #ffffff;
  --win11-text-primary: #191919;
  --win11-text-secondary: #606060;
  --win11-accent: #0067c0;
  --win11-inner-bg: rgba(0, 0, 0, 0.03);

  background-color: var(--bg-color);
  border: 1px solid color-mix(in srgb, var(--font-color), transparent 80%);
  border-radius: 8px;
  padding: 20px 24px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  color: var(--font-color);
  display: flex;
  flex-direction: column;
  gap: 12px;
  transition:
    background-color 0.15s ease,
    border-color 0.15s ease;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .page-subtitle {
    color: #a6a6a6;
  }
  .about-card {
    --win11-border-color: rgba(255, 255, 255, 0.08);
    --win11-bg: #1c1c1c;
    --win11-text-primary: #ffffff;
    --win11-text-secondary: #a6a6a6;
    --win11-accent: #60cdff;
    --win11-inner-bg: rgba(255, 255, 255, 0.04);
  }
}

.card-title {
  font-size: 15px;
  font-weight: 600;
  margin-bottom: 8px;
}

/* 产品信息 */
.product-card {
  flex-direction: row;
  align-items: center;
  gap: 24px;
  padding: 24px;
}

.product-logo {
  width: 72px;
  height: 72px;
  object-fit: contain;
  flex-shrink: 0;
}

.product-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.product-title {
  font-size: 22px;
  font-weight: 600;
}

.product-version {
  font-size: 13px;
  font-weight: 500;
}

.product-desc {
  font-size: 13px;
  margin-top: 4px;
}

/* 开发者信息 */
.dev-columns {
  display: flex;
  gap: 16px;
}

.dev-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
  background-color: var(--win11-inner-bg);
  padding: 14px 16px;
  border-radius: 6px;
  border: 1px solid color-mix(in srgb, var(--font-color), transparent 80%);
}

.col-label {
  font-size: 12px;
}

.col-value {
  font-size: 14px;
  font-weight: 500;
}

/* 官网与支持 */
.support-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 16px;
  border-radius: 6px;
  border: 1px solid color-mix(in srgb, var(--font-color), transparent 80%);
}

.support-text {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.support-title {
  font-size: 14px;
  font-weight: 600;
}

.support-desc {
  font-size: 13px;
}

.support-link {
  font-size: 14px;
  font-weight: 500;
  text-decoration: none;
  padding: 6px 12px;
  border-radius: 4px;
  color: var(--font-color);
  transition: background-color 0.2s;
}

.support-link:hover {
  background-color: rgba(120, 120, 120, 0.15);
}
</style>
