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
                            <li><a class="dropdown-item" href="https://cv.pjmcode.top/">开发者信息</a></li>
                            <li><a class="dropdown-item" href="https://github.com/pjimming/baize">开源代码</a></li>
                            <li><a class="dropdown-item" href="#">赞助</a></li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script></script>

<style scoped></style>
