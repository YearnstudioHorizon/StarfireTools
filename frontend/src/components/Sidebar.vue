<script setup>
import { computed, ref } from "vue";
import { useRoute, useRouter } from "vue-router";

const expanded = ref(false);
const router = useRouter();
const route = useRoute();

const itemsTop = [
  { key: "home", label: "主页", path: "/" },
  { key: "list", label: "功能", path: "/Features" },
  { key: "help", label: "关于", path: "/About" },
];

const activeNavIndex = computed(() =>
  itemsTop.findIndex((item) => item.path === route.path),
);

const navSelectionStyle = computed(() => {
  const itemHeight = 54;
  const gap = 8;

  return {
    "--selection-y": `${activeNavIndex.value * (itemHeight + gap)}px`,
  };
});

const onSelect = (path) => {
  if (route.path === path) return;
  router.push(path);
};
</script>

<template>
  <aside
    class="sf-sidebar"
    :class="{ 'is-expanded': expanded }"
    aria-label="Sidebar"
  >
    <button
      class="sf-sidebar__iconbtn sf-sidebar__menu"
      type="button"
      aria-label="Menu"
      data-wails-no-drag
      @click="expanded = !expanded"
    >
      <svg viewBox="0 0 24 24" aria-hidden="true">
        <path
          d="M4 6.5h16M4 12h16M4 17.5h16"
          fill="none"
          stroke="var(--font-color)"
          stroke-width="1.8"
          stroke-linecap="round"
        />
      </svg>
    </button>

    <nav
      class="sf-sidebar__nav"
      :style="navSelectionStyle"
      aria-label="Primary"
    >
      <div
        class="sf-sidebar__selection"
        :style="{ opacity: activeNavIndex >= 0 ? 1 : 0 }"
        aria-hidden="true"
      >
        <span class="sf-sidebar__selection-indicator"></span>
      </div>

      <button
        v-for="it in itemsTop"
        :key="it.key"
        class="sf-sidebar__item"
        type="button"
        :aria-label="it.label"
        :aria-current="route.path === it.path ? 'page' : undefined"
        data-wails-no-drag
        @click="onSelect(it.path)"
      >
        <span class="sf-sidebar__item-inner">
          <span class="sf-sidebar__icon" aria-hidden="true">
            <svg v-if="it.key === 'home'" viewBox="0 0 24 24">
              <path
                d="M4 11.5L12 5l8 6.5V20a1.5 1.5 0 0 1-1.5 1.5H5.5A1.5 1.5 0 0 1 4 20v-8.5Z"
                fill="none"
                stroke="var(--font-color)"
                stroke-width="1.8"
                stroke-linejoin="round"
              />
            </svg>
            <svg v-else-if="it.key === 'list'" viewBox="0 0 24 24">
              <path
                d="M9 6h12M9 12h12M9 18h12"
                fill="none"
                stroke="var(--font-color)"
                stroke-width="1.8"
                stroke-linecap="round"
              />
              <path
                d="M4.5 6h.01M4.5 12h.01M4.5 18h.01"
                fill="none"
                stroke="var(--font-color)"
                stroke-width="3"
                stroke-linecap="round"
              />
            </svg>
            <svg v-else viewBox="0 0 24 24">
              <path
                d="M9.2 9a3 3 0 1 1 4.9 2.2c-.9.7-1.6 1.3-1.6 2.8v.2"
                fill="none"
                stroke="var(--font-color)"
                stroke-width="1.8"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
              <path
                d="M12 18.2h.01"
                fill="none"
                stroke="var(--font-color)"
                stroke-width="3"
                stroke-linecap="round"
              />
            </svg>
          </span>
          <span class="sf-sidebar__label">{{ it.label }}</span>
        </span>
      </button>
    </nav>

    <div class="sf-sidebar__spacer" aria-hidden="true"></div>

    <button
      class="sf-sidebar__item sf-sidebar__item--standalone"
      :class="{ 'is-active': route.path === '/settings' }"
      type="button"
      aria-label="Settings"
      :aria-current="route.path === '/settings' ? 'page' : undefined"
      data-wails-no-drag
      @click="onSelect('/settings')"
    >
      <span class="sf-sidebar__item-inner">
        <span class="sf-sidebar__icon" aria-hidden="true">
          <svg viewBox="0 0 24 24">
            <path
              d="M12 15.2a3.2 3.2 0 1 0 0-6.4 3.2 3.2 0 0 0 0 6.4Z"
              fill="none"
              stroke="var(--font-color)"
              stroke-width="1.8"
            />
            <path
              d="M19.3 13.1v-2.2l-2-.4a7.5 7.5 0 0 0-.7-1.6l1.2-1.7-1.5-1.5-1.7 1.2c-.5-.3-1-.5-1.6-.7L13.1 4.7h-2.2l-.4 2a7.5 7.5 0 0 0-1.6.7L7.2 6.2 5.7 7.7l1.2 1.7c-.3.5-.5 1-.7 1.6l-2 .4v2.2l2 .4c.2.6.4 1.1.7 1.6l-1.2 1.7 1.5 1.5 1.7-1.2c.5.3 1 .5 1.6.7l.4 2h2.2l.4-2c.6-.2 1.1-.4 1.6-.7l1.7 1.2 1.5-1.5-1.2-1.7c.3-.5.5-1 .7-1.6l2-.4Z"
              fill="none"
              stroke="var(--font-color)"
              stroke-width="1.4"
              stroke-linejoin="round"
            />
          </svg>
        </span>
        <span class="sf-sidebar__label">设置</span>
      </span>
    </button>
  </aside>
</template>

<style scoped>
.sf-sidebar {
  width: 80px;
  height: 100%;
  flex: 0 0 auto;
  overflow: hidden;
  background: var(--bg-color);
  border-right: 1px solid color-mix(in srgb, var(--font-color), transparent 90%);
  box-shadow:
    0 0 0 1px rgba(0, 0, 0, 0.03),
    0 10px 30px rgba(0, 0, 0, 0.08);
  display: flex;
  flex-direction: column;
  align-items: stretch;
  padding: 10px 8px;
  box-sizing: border-box;
  transition: width 0.22s cubic-bezier(0.32, 0.72, 0, 1);
}

.sf-sidebar.is-expanded {
  width: 220px;
}

.sf-sidebar__iconbtn {
  width: 100%;
  height: 44px;
  margin: 0;
  border: 0;
  border-radius: 10px;
  background: transparent;
  color: var(--font-color);
  display: flex;
  align-items: center;
  justify-content: flex-start;
  cursor: pointer;
  flex: 0 0 auto;
  box-sizing: border-box;
  padding: 0 12px 0 21px;
  transition: background-color 0.12s ease;
}

.sf-sidebar__iconbtn:hover {
  background: color-mix(in srgb, var(--font-color), transparent 94%);
}

.sf-sidebar__iconbtn svg {
  width: 22px;
  height: 22px;
}

.sf-sidebar__nav {
  position: relative;
  width: 100%;
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.sf-sidebar__selection {
  position: absolute;
  top: 0;
  left: 8px;
  width: 50px;
  height: 54px;
  border-radius: 10px;
  background: color-mix(in srgb, var(--font-color), transparent 90%);
  transform: translateY(var(--selection-y, 0));
  transition:
    transform 0.22s cubic-bezier(0.32, 0.72, 0, 1),
    width 0.22s cubic-bezier(0.32, 0.72, 0, 1),
    left 0.22s cubic-bezier(0.32, 0.72, 0, 1),
    opacity 0.14s ease;
  pointer-events: none;
}

.sf-sidebar.is-expanded .sf-sidebar__selection {
  left: 8px;
  width: calc(100% - 8px);
  height: 54px;
  border-radius: 10px;
  transform: translateY(var(--selection-y, 0));
}

.sf-sidebar__selection-indicator {
  position: absolute;
  left: -8px;
  top: 50%;
  width: 5px;
  height: 24px;
  border-radius: 999px;
  background: #0a6bff;
  transform: translateY(-50%);
}

.sf-sidebar__item {
  width: 100%;
  padding: 0;
  border: 0;
  background: transparent;
  color: var(--font-color);
  cursor: pointer;
  position: relative;
  z-index: 1;
}

.sf-sidebar__item-inner {
  width: 100%;
  height: 54px;
  margin: 0;
  display: flex;
  align-items: center;
  box-sizing: border-box;
  overflow: hidden;
  padding: 0 12px 0 21px;
}

.sf-sidebar__item:hover .sf-sidebar__item-inner {
  background: color-mix(in srgb, var(--font-color), transparent 95%);
  border-radius: 10px;
}

.sf-sidebar__icon {
  flex: 0 0 24px;
  width: 24px;
  height: 24px;
  display: grid;
  place-items: center;
}

.sf-sidebar__icon svg {
  width: 24px;
  height: 24px;
}

.sf-sidebar__label {
  display: block;
  max-width: 0;
  flex: 0 0 auto;
  overflow: hidden;
  margin-left: 0;
  line-height: 1;
  font-size: 13px;
  font-weight: 600;
  color: var(--font-color);
  white-space: nowrap;
  opacity: 0;
  transform: translateX(-6px);
  transition:
    max-width 0.22s cubic-bezier(0.32, 0.72, 0, 1),
    margin-left 0.22s cubic-bezier(0.32, 0.72, 0, 1),
    opacity 0.16s ease,
    transform 0.22s cubic-bezier(0.32, 0.72, 0, 1);
  pointer-events: none;
}

.sf-sidebar.is-expanded .sf-sidebar__label {
  max-width: 120px;
  margin-left: 10px;
  opacity: 1;
  transform: translateX(0);
}

.sf-sidebar__spacer {
  flex: 1;
  min-height: 12px;
}

.sf-sidebar__item--standalone {
  flex: 0 0 auto;
}

.sf-sidebar__item--standalone::before {
  content: "";
  position: absolute;
  inset: 0;
  width: 50px;
  height: 54px;
  margin: 0 auto;
  border-radius: 10px;
  background: color-mix(in srgb, var(--font-color), transparent 90%);
  opacity: 0;
  transition:
    opacity 0.14s ease,
    width 0.22s cubic-bezier(0.32, 0.72, 0, 1),
    height 0.22s cubic-bezier(0.32, 0.72, 0, 1);
  pointer-events: none;
}

.sf-sidebar.is-expanded .sf-sidebar__item--standalone::before {
  width: 100%;
  height: 54px;
}

.sf-sidebar__item--standalone.is-active::before {
  opacity: 1;
}

@media (prefers-reduced-motion: reduce) {
  .sf-sidebar,
  .sf-sidebar__iconbtn,
  .sf-sidebar__selection,
  .sf-sidebar__item-inner,
  .sf-sidebar__label,
  .sf-sidebar__item--standalone::before {
    transition: none;
  }
}
</style>
