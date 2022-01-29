<template>
    <v-container>
        <v-card height="600px" width="1000px">
                    <v-simple-table>
                        <template>
                            <thead>
                                <tr>
                                    <th>
                                        更新日付
                                    </th>
                                    <th>
                                        更新内容
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td>2021/11/25</td>
                                    <td>テスト</td>
                                </tr>
                            </tbody>
                        </template>
                    </v-simple-table>
                </v-card>
    </v-container>
</template>

<script>
export default ({
    name: 'announce'
})
</script>
