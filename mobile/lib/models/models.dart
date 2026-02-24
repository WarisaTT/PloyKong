// lib/models/models.dart

class UserModel {
  final String id;
  final String email;
  final String name;
  final String? avatarUrl;
  final String? bio;
  final String plan;

  UserModel({
    required this.id,
    required this.email,
    required this.name,
    this.avatarUrl,
    this.bio,
    required this.plan,
  });

  factory UserModel.fromJson(Map<String, dynamic> j) => UserModel(
        id: j['id'],
        email: j['email'],
        name: j['name'],
        avatarUrl: j['avatar_url'],
        bio: j['bio'],
        plan: j['plan'] ?? 'free',
      );
}

class PortfolioModel {
  final String id;
  final String slug;
  final String title;
  final bool isPublished;
  final int viewCount;
  final DateTime updatedAt;

  PortfolioModel({
    required this.id,
    required this.slug,
    required this.title,
    required this.isPublished,
    required this.viewCount,
    required this.updatedAt,
  });

  factory PortfolioModel.fromJson(Map<String, dynamic> j) => PortfolioModel(
        id: j['id'],
        slug: j['slug'],
        title: j['title'],
        isPublished: j['is_published'] ?? false,
        viewCount: j['view_count'] ?? 0,
        updatedAt: DateTime.parse(j['updated_at']),
      );
}

class AnalyticsSummary {
  final int totalViews;
  final int todayViews;
  final int weekViews;
  final int pdfDownloads;
  final int hireClicks;
  final int aiChatCount;
  final List<DailyView> dailyViews;
  final List<CountryStat> countryStats;
  final List<SourceStat> sourceStats;

  AnalyticsSummary({
    required this.totalViews,
    required this.todayViews,
    required this.weekViews,
    required this.pdfDownloads,
    required this.hireClicks,
    required this.aiChatCount,
    required this.dailyViews,
    required this.countryStats,
    required this.sourceStats,
  });

  factory AnalyticsSummary.fromJson(Map<String, dynamic> j) => AnalyticsSummary(
        totalViews: j['total_views'] ?? 0,
        todayViews: j['today_views'] ?? 0,
        weekViews: j['week_views'] ?? 0,
        pdfDownloads: j['pdf_downloads'] ?? 0,
        hireClicks: j['hire_clicks'] ?? 0,
        aiChatCount: j['ai_chat_count'] ?? 0,
        dailyViews: (j['daily_views'] as List? ?? [])
            .map((e) => DailyView.fromJson(e))
            .toList(),
        countryStats: (j['country_stats'] as List? ?? [])
            .map((e) => CountryStat.fromJson(e))
            .toList(),
        sourceStats: (j['source_stats'] as List? ?? [])
            .map((e) => SourceStat.fromJson(e))
            .toList(),
      );
}

class DailyView {
  final String date;
  final int count;
  DailyView({required this.date, required this.count});
  factory DailyView.fromJson(Map<String, dynamic> j) =>
      DailyView(date: j['date'], count: j['count'] ?? 0);
}

class CountryStat {
  final String countryCode;
  final int count;
  CountryStat({required this.countryCode, required this.count});
  factory CountryStat.fromJson(Map<String, dynamic> j) =>
      CountryStat(countryCode: j['country_code'] ?? '--', count: j['count'] ?? 0);
}

class SourceStat {
  final String source;
  final int count;
  SourceStat({required this.source, required this.count});
  factory SourceStat.fromJson(Map<String, dynamic> j) =>
      SourceStat(source: j['source'] ?? 'Direct', count: j['count'] ?? 0);
}

class NotificationModel {
  final String id;
  final String type; // 'pdf_download' | 'hire_click' | 'view'
  final String portfolioTitle;
  final String? country;
  final DateTime createdAt;
  bool isRead;

  NotificationModel({
    required this.id,
    required this.type,
    required this.portfolioTitle,
    this.country,
    required this.createdAt,
    this.isRead = false,
  });

  factory NotificationModel.fromJson(Map<String, dynamic> j) =>
      NotificationModel(
        id: j['id'],
        type: j['type'],
        portfolioTitle: j['portfolio_title'] ?? 'My Portfolio',
        country: j['country'],
        createdAt: DateTime.parse(j['created_at']),
        isRead: j['is_read'] ?? false,
      );

  String get emoji {
    switch (type) {
      case 'pdf_download':
        return '📄';
      case 'hire_click':
        return '🎯';
      case 'view':
        return '👁️';
      default:
        return '🔔';
    }
  }

  String get title {
    switch (type) {
      case 'pdf_download':
        return 'มีคนดาวน์โหลด Resume ของคุณ!';
      case 'hire_click':
        return 'มีคนกดปุ่ม "Hire Me"!';
      case 'view':
        return 'มีคนเข้าชมพอร์ตของคุณ';
      default:
        return 'การแจ้งเตือนใหม่';
    }
  }
}
