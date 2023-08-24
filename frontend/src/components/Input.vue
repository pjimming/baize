<template>
    <div class="input-group input-group-lg">
        <span class="input-group-text" id="inputGroup-sizing-default">请输入项目的绝对路径</span>
        <input type="text" class="form-control" placeholder="例如：/Users/panjiangming/Project/baize/backend"
            v-model="inputValue">
        <button class="btn btn-primary" @click="getModulePath">Submit</button>
    </div>
</template>

<script>
import { fetchModuleInfo } from '@/services/api';
import store from '@/store';

// const data = {
//     modulePath: "",
//     dir: "",
//     projectInfo: {
//         modulePath: "",
//         moduleName: "github.com/pjimming/baize",
//         moduleVersion: "go 1.18",
//         otherPkgList: [{
//             "name": "github.com/stretchr/testify",
//             "version": "v1.8.4"
//         },
//         {
//             "name": "github.com/zeromicro/go-zero",
//             "version": "v1.5.4"
//         }
//         ],
//         otherPkgCount: 2,
//         projectPkgList: [
//             "github.com/pjimming/baize/common/errorx",
//             "github.com/pjimming/baize/common/httpresp",
//             "github.com/pjimming/baize/core",
//             "github.com/pjimming/baize/core/helper",
//             "github.com/pjimming/baize/core/internal/config",
//             "github.com/pjimming/baize/core/internal/handler",
//             "github.com/pjimming/baize/core/internal/handler/basic",
//             "github.com/pjimming/baize/core/internal/handler/local",
//             "github.com/pjimming/baize/core/internal/logic/basic",
//             "github.com/pjimming/baize/core/internal/logic/local",
//             "github.com/pjimming/baize/core/internal/svc",
//             "github.com/pjimming/baize/core/internal/types"
//         ],
//         projectPkgCount: 12,
//         goFileList: [{
//             "name": "D:\\GoProject\\baize\\backend\\common\\errorx\\codeerror.go",
//             "size": 1023
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\common\\httpresp\\response.go",
//             "size": 1259
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\core.go",
//             "size": 807
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\helper\\helper.go",
//             "size": 4676
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\helper\\helper_test.go",
//             "size": 2730
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\config\\config.go",
//             "size": 98
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\handler\\basic\\pinghandler.go",
//             "size": 318
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\handler\\local\\generategraphhandler.go",
//             "size": 792
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\handler\\local\\getgofileshandler.go",
//             "size": 783
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\handler\\local\\getmoduleinfohandler.go",
//             "size": 792
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\handler\\local\\getpackageshandler.go",
//             "size": 786
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\handler\\routes.go",
//             "size": 1140
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\logic\\basic\\pinglogic.go",
//             "size": 523
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\logic\\local\\generategraphlogic.go",
//             "size": 2400
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\logic\\local\\getgofileslogic.go",
//             "size": 924
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\logic\\local\\getmoduleinfologic.go",
//             "size": 1268
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\logic\\local\\getpackageslogic.go",
//             "size": 1162
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\svc\\servicecontext.go",
//             "size": 230
//         },
//         {
//             "name": "D:\\GoProject\\baize\\backend\\core\\internal\\types\\types.go",
//             "size": 1240
//         }
//         ],
//         goFileCount: 19,
//     },
//     graphData: {
//         nodes: [{
//             "id": "github.com/pjimming/baize/common/errorx",
//             "label": "github.com/pjimming/baize/common/errorx"
//         },
//         {
//             "id": "github.com/pjimming/baize/common/httpresp",
//             "label": "github.com/pjimming/baize/common/httpresp"
//         },
//         {
//             "id": "github.com/pjimming/baize/core",
//             "label": "github.com/pjimming/baize/core"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/helper",
//             "label": "github.com/pjimming/baize/core/helper"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/config",
//             "label": "github.com/pjimming/baize/core/internal/config"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/handler",
//             "label": "github.com/pjimming/baize/core/internal/handler"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/handler/basic",
//             "label": "github.com/pjimming/baize/core/internal/handler/basic"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/handler/local",
//             "label": "github.com/pjimming/baize/core/internal/handler/local"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/logic/basic",
//             "label": "github.com/pjimming/baize/core/internal/logic/basic"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/logic/local",
//             "label": "github.com/pjimming/baize/core/internal/logic/local"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/svc",
//             "label": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "id": "github.com/pjimming/baize/core/internal/types",
//             "label": "github.com/pjimming/baize/core/internal/types"
//         }
//         ],
//         edges: [{
//             "source": "github.com/pjimming/baize/core/internal/handler/local",
//             "target": "github.com/pjimming/baize/core/internal/types"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/logic/local",
//             "target": "github.com/pjimming/baize/common/errorx"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/logic/local",
//             "target": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "source": "github.com/pjimming/baize/common/httpresp",
//             "target": "github.com/pjimming/baize/common/errorx"
//         },
//         {
//             "source": "github.com/pjimming/baize/core",
//             "target": "github.com/pjimming/baize/core/internal/config"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler/basic",
//             "target": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler/local",
//             "target": "github.com/pjimming/baize/common/errorx"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler/local",
//             "target": "github.com/pjimming/baize/common/httpresp"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler",
//             "target": "github.com/pjimming/baize/core/internal/handler/local"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/logic/local",
//             "target": "github.com/pjimming/baize/core/helper"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/helper",
//             "target": "github.com/pjimming/baize/core/internal/types"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler/local",
//             "target": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler",
//             "target": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/logic/basic",
//             "target": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/logic/local",
//             "target": "github.com/pjimming/baize/core/internal/types"
//         },
//         {
//             "source": "github.com/pjimming/baize/core",
//             "target": "github.com/pjimming/baize/core/internal/handler"
//         },
//         {
//             "source": "github.com/pjimming/baize/core",
//             "target": "github.com/pjimming/baize/core/internal/svc"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler/local",
//             "target": "github.com/pjimming/baize/core/internal/logic/local"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/handler",
//             "target": "github.com/pjimming/baize/core/internal/handler/basic"
//         },
//         {
//             "source": "github.com/pjimming/baize/core/internal/svc",
//             "target": "github.com/pjimming/baize/core/internal/config"
//         }
//         ]
//     }
// };

export default {
    name: "InputVue",
    data() {
        return {
            inputValue: "",
        };
    },
    methods: {
        async getModulePath() {
            // store.replaceState(data);

            const queryParams = {
                dir: this.inputValue,
            }

            store.commit('setDir', this.inputValue);

            try {
                await fetchModuleInfo(queryParams);
            } catch (error) {
                console.error(error.message);
            }
        },
    },
}
</script>

<style scoped></style>
