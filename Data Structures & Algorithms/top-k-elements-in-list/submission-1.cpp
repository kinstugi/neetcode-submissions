class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        auto comparePairs = [](const std::pair<int, int>& a, const std::pair<int, int>& b) {
        return a.first > b.first; // For min-heap, return true if 'a' has a larger first element
        };
        map<int, int> mp;
        for (int num: nums)
        {
            mp[num]++;
        }
        std::priority_queue<std::pair<int, int>, std::vector<std::pair<int, int>>, decltype(comparePairs)> min_pq(comparePairs);
        
        for (const auto &[num, cnt]: mp)
        {
            min_pq.push({cnt, num});
            if (min_pq.size() > k)
            {
                min_pq.pop();
            }
        }

        std::vector<int> ans;
        ans.reserve(10);
        while(!min_pq.empty())
        {
            pair<int, int> pr = min_pq.top();
            min_pq.pop();
            ans.push_back(pr.second);
        }
        return ans;
    }
};
