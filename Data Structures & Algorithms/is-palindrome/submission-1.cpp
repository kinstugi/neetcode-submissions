class Solution {
public:
    bool isPalindrome(string s) {
        std::vector<char> nStr;
        nStr.reserve(100);
        for (char ch: s)
        {
            char c = tolower(ch);
            if (!isalnum(c)) continue;
            nStr.push_back(c);
        }
        
        int l = 0, r = nStr.size()-1;
        while(l <= r){
            if (nStr[l] != nStr[r]) return false;
            l++;
            r--;
        }
        return true;
    }
};
