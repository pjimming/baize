<template>
    <nav class="navbar navbar-expand-lg fixed-top bg-body-tertiary bg-primary" data-bs-theme="dark">
        <div class="container-fluid">
            <router-link class="navbar-brand" :to="{ name: 'home' }">白泽</router-link>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown"
                aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNavDropdown">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <router-link class="nav-link" aria-current="page" :to="{ name: 'home' }">首页</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" :to="{ name: 'about' }">关于白泽</router-link>
                    </li>
                    <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
                            aria-expanded="false">
                            其他
                        </a>
                        <ul class="dropdown-menu">
                            <li><a class="dropdown-item" href="https://cv.pjmcode.top/">开发者信息
                                    <svg preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24" width="1.2em"
                                        height="1.2em" class="link-icon">
                                        <path fill="currentColor"
                                            d="M10 6v2H5v11h11v-5h2v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h6zm11-3v8h-2V6.413l-7.793 7.794l-1.414-1.414L17.585 5H13V3h8z">
                                        </path>
                                    </svg>
                                </a></li>
                            <li><a class="dropdown-item" href="https://github.com/pjimming/baize">开源代码
                                    <svg preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24" width="1.2em"
                                        height="1.2em" class="link-icon">
                                        <path fill="currentColor"
                                            d="M10 6v2H5v11h11v-5h2v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h6zm11-3v8h-2V6.413l-7.793 7.794l-1.414-1.414L17.585 5H13V3h8z">
                                        </path>
                                    </svg>
                                </a></li>
                            <li><a class="dropdown-item" href="#">赞助</a></li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script>
export default {
    name: 'NavBar',
}
</script>

<style scoped></style>
